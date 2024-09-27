/**
 * Copyright 2019 The Gamma Authors.
 *
 * This source code is licensed under the Apache License, Version 2.0 license
 * found in the LICENSE file in the root directory of this source tree.
 */

#include "rocksdb_raw_vector.h"

#include <stdio.h>

#include "rocksdb/table.h"
#include "util/log.h"
#include "util/utils.h"

namespace vearch {

RocksDBRawVector::RocksDBRawVector(VectorMetaInfo *meta_info,
                                   const StoreParams &store_params,
                                   bitmap::BitmapManager *docids_bitmap,
                                   StorageManager *storage_mgr, int cf_id)
    : RawVector(meta_info, docids_bitmap, store_params) {
  storage_mgr_ = storage_mgr;
  cf_id_ = cf_id;
}

RocksDBRawVector::~RocksDBRawVector() {}

int RocksDBRawVector::GetDiskVecNum(int &vec_num) {
  if (vec_num <= 0) return 0;
  int max_id_in_disk = vec_num - 1;
  for (int i = max_id_in_disk; i >= 0; --i) {
    auto result = storage_mgr_->Get(cf_id_, i);
    if (result.first.ok()) {
      vec_num = i + 1;
      LOG(INFO) << "In the disk rocksdb vec_num=" << vec_num;
      return 0;
    }
  }
  vec_num = 0;
  LOG(INFO) << "In the disk rocksdb vec_num=" << vec_num;
  return 0;
}

Status RocksDBRawVector::Load(int vec_num) {
  if (vec_num == 0) return Status::OK();
  std::string key, value;
  ToRowKey(vec_num - 1, key);
  rocksdb::Status s = storage_mgr_->db_->Get(rocksdb::ReadOptions(),
                                             storage_mgr_->cf_handles_[cf_id_],
                                             rocksdb::Slice(key), &value);
  if (!s.ok()) {
    std::string msg = std::string("load vectors, get error:") + s.ToString() +
                      ", expected key=" + key;
    LOG(ERROR) << msg;
    return Status::IOError(msg);
  }
  MetaInfo()->size_ = vec_num;
  LOG(INFO) << "rocksdb load success! vec_num=" << vec_num;
  return Status::OK();
}

int RocksDBRawVector::InitStore(std::string &vec_name) {
  block_cache_size_ = (size_t)store_params_.cache_size * 1024 * 1024;

  // std::shared_ptr<rocksdb::Cache> cache =
  //     rocksdb::NewLRUCache(block_cache_size_);
  // table_options_.block_cache = cache;
  // rocksdb::Options options;
  // options.table_factory.reset(NewBlockBasedTableFactory(table_options_));

  // options.IncreaseParallelism();
  // // options.OptimizeLevelStyleCompaction();
  // // create the DB if it's not already present
  // options.create_if_missing = true;

  // std::string db_path = this->root_path_ + "/" + meta_info_->Name();
  // if (!utils::isFolderExist(db_path.c_str())) {
  //   mkdir(db_path.c_str(), S_IRWXU | S_IRWXG | S_IROTH | S_IXOTH);
  // }

  // // open DB
  // rocksdb::Status s = rocksdb::DB::Open(options, db_path, &db_);
  // if (!s.ok()) {
  //   LOG(ERROR) << "open rocks db error: " << s.ToString();
  //   return -1;
  // }
  // LOG(INFO) << "rocks raw vector init success! name=" << meta_info_->Name()
  //           << ", block cache size=" << block_cache_size_ << "Bytes";

  return 0;
}

int RocksDBRawVector::GetVector(long vid, const uint8_t *&vec,
                                bool &deletable) const {
  if ((size_t)vid >= meta_info_->Size() || vid < 0) {
    return 1;
  }
  auto result = storage_mgr_->Get(cf_id_, vid);
  if (!result.first.ok()) {
    LOG(ERROR) << "rocksdb get error:" << result.first.ToString()
               << ", vid=" << vid;
    return result.first.code();
  }
  vec = new uint8_t[vector_byte_size_];
  assert((size_t)vector_byte_size_ == result.second.size());
  memcpy((void *)vec, result.second.c_str(), vector_byte_size_);

  deletable = true;
  return 0;
}

int RocksDBRawVector::Gets(const std::vector<int64_t> &vids,
                           ScopeVectors &vecs) const {
  size_t k = vids.size();
  std::vector<std::string> keys_data(k);
  std::vector<rocksdb::Slice> keys;
  std::vector<rocksdb::ColumnFamilyHandle *> column_families;
  keys.reserve(k);
  column_families.reserve(k);

  size_t j = 0;
  for (size_t i = 0; i < k; i++) {
    if (vids[i] < 0) {
      continue;
    }
    ToRowKey((int)vids[i], keys_data[i]);
    keys.emplace_back(std::move(keys_data[i]));
    column_families.emplace_back(storage_mgr_->cf_handles_[cf_id_]);
    ++j;
    // LOG(INFO) << "i=" << i << "key=" << keys[i].ToString();
  }

  std::vector<std::string> values(j);
  std::vector<rocksdb::Status> statuses = storage_mgr_->db_->MultiGet(
      rocksdb::ReadOptions(), column_families, keys, &values);
  assert(statuses.size() == j);

  j = 0;
  for (size_t i = 0; i < k; ++i) {
    if (vids[i] < 0) {
      vecs.Add(nullptr, true);
      continue;
    }
    if (!statuses[j].ok()) {
      LOG(ERROR) << "rocksdb multiget error:" << statuses[j].ToString()
                 << ", key=" << keys[j].ToString();
      return 2;
    }
    uint8_t *vector = new uint8_t[vector_byte_size_];
    assert((size_t)vector_byte_size_ == values[j].size());
    memcpy(vector, values[j].c_str(), vector_byte_size_);

    vecs.Add(vector, true);
    ++j;
  }
  return 0;
}

int RocksDBRawVector::AddToStore(uint8_t *v, int len) {
  return UpdateToStore(meta_info_->Size(), v, len);
}

size_t RocksDBRawVector::GetStoreMemUsage() {
  //   size_t cache_mem = table_options_.block_cache->GetUsage();
  //   std::string index_mem;
  //   db_->GetProperty("rocksdb.estimate-table-readers-mem", &index_mem);
  //   std::string memtable_mem;
  //   db_->GetProperty("rocksdb.cur-size-all-mem-tables", &memtable_mem);
  //   size_t pin_mem = table_options_.block_cache->GetPinnedUsage();
  // #ifdef DEBUG
  //   LOG(INFO) << "rocksdb mem usage: block cache=" << cache_mem
  //             << ", index and filter=" << index_mem
  //             << ", memtable=" << memtable_mem
  //             << ", iterators pinned=" << pin_mem;
  // #endif
  //   return cache_mem + pin_mem;
  return 0;
}

int RocksDBRawVector::UpdateToStore(int vid, uint8_t *v, int len) {
  if (v == nullptr || len != meta_info_->Dimension() * meta_info_->DataSize())
    return -1;

  std::string key;
  ToRowKey(vid, key);
  rocksdb::Status s = storage_mgr_->db_->Put(
      rocksdb::WriteOptions(), storage_mgr_->cf_handles_[cf_id_],
      rocksdb::Slice(key),
      rocksdb::Slice((const char *)v, this->vector_byte_size_));
  if (!s.ok()) {
    LOG(ERROR) << "rocksdb update error:" << s.ToString() << ", key=" << key;
    return -1;
  }
  return 0;
}

int RocksDBRawVector::GetVectorHeader(int start, int n, ScopeVectors &vecs,
                                      std::vector<int> &lens) {
  if (start < 0 || (size_t)start + n > meta_info_->Size()) {
    return -1;
  }

  rocksdb::Iterator *it = storage_mgr_->db_->NewIterator(
      rocksdb::ReadOptions(), storage_mgr_->cf_handles_[cf_id_]);
  std::string start_key, end_key;
  ToRowKey(start, start_key);
  ToRowKey(start + n, end_key);
  it->Seek(rocksdb::Slice(start_key));
  int dimension = meta_info_->Dimension();
  uint8_t *vectors = new uint8_t[(uint64_t)dimension * n * data_size_];
  uint8_t *dst = vectors;
  for (int c = 0; c < n; c++, it->Next()) {
    if (!it->Valid()) {
      LOG(ERROR) << "rocksdb iterator error, vid=" << start + c;
      delete it;
      delete[] vectors;
      return -1;
    }

    rocksdb::Slice value = it->value();
    std::string vstr = value.ToString();
    assert((size_t)vector_byte_size_ == vstr.size());
    memcpy(dst, vstr.c_str(), vector_byte_size_);

#ifdef DEBUG
    std::string expect_key;
    ToRowKey(c + start, expect_key);
    std::string key = it->key().ToString();
    if (key != expect_key) {
      LOG(ERROR) << "vid=" << c + start << ", invalid key=" << key
                 << ", expect=" << expect_key;
    }
#endif
    dst += (size_t)dimension * data_size_;
  }
  delete it;
  vecs.Add(vectors);
  lens.push_back(n);
  return 0;
}

void RocksDBRawVector::ToRowKey(int vid, std::string &key) const {
  char data[11];
  snprintf(data, 11, "%010d", vid);
  key.assign(data, 10);
}

}  // namespace vearch
