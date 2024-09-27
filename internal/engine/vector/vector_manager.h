
/**
 * Copyright 2019 The Gamma Authors.
 *
 * This source code is licensed under the Apache License, Version 2.0 license
 * found in the LICENSE file in the root directory of this source tree.
 */

#pragma once

#include <map>
#include <string>

#include "common/gamma_common_data.h"
#include "index/index_model.h"
#include "util/bitmap_manager.h"
#include "util/log.h"
#include "util/status.h"
#include "vector/raw_vector.h"

namespace vearch {

class VectorManager {
 public:
  VectorManager(const VectorStorageType &store_type,
                bitmap::BitmapManager *docids_bitmap,
                const std::string &root_path);
  ~VectorManager();

  Status SetVectorStoreType(std::string &index_type,
                            std::string &store_type_str,
                            VectorStorageType &store_type);

  Status CreateRawVector(struct VectorInfo &vector_info,
                         std::string &index_type,
                         std::map<std::string, int> &vec_dups, TableInfo &table,
                         RawVector **vec, int cf_id,
                         StorageManager *storage_mgr);

  void DestroyRawVectors();

  Status CreateVectorIndex(std::string &index_type, std::string &index_params,
                           RawVector *vec, int training_threshold,
                           bool destroy_vec,
                           std::map<std::string, IndexModel *> &vector_indexes);

  void DestroyVectorIndexes();

  void DescribeVectorIndexes();

  Status CreateVectorIndexes(
      int training_threshold,
      std::map<std::string, IndexModel *> &vector_indexes);

  void SetVectorIndexes(
      std::map<std::string, IndexModel *> &rebuild_vector_indexes);

  Status CreateVectorTable(TableInfo &table, std::vector<int> &vector_cf_ids,
                           StorageManager *storage_mgr);

  int AddToStore(int docid,
                 std::unordered_map<std::string, struct Field> &fields);

  int Update(int docid, std::unordered_map<std::string, struct Field> &fields);

  int TrainIndex(std::map<std::string, IndexModel *> &vector_indexes);

  int AddRTVecsToIndex(bool &index_is_dirty);

  // int Add(int docid, const std::vector<Field *> &field_vecs);
  Status Search(GammaQuery &query, GammaResult *results);

  int GetVector(const std::vector<std::pair<std::string, int>> &fields_ids,
                std::vector<std::string> &vec, bool is_bytearray = false);

  int GetDocVector(int docid, std::string &field_name,
                   std::vector<uint8_t> &vec);

  void GetTotalMemBytes(long &index_total_mem_bytes,
                        long &vector_total_mem_bytes);

  int Dump(const std::string &path, int dump_docid, int max_docid);
  int Load(const std::vector<std::string> &path, int &doc_num);

  bool Contains(std::string &field_name);

  void VectorNames(std::vector<std::string> &names) {
    for (const auto &it : raw_vectors_) {
      names.push_back(it.first);
    }
  }

  std::map<std::string, IndexModel *> &VectorIndexes() {
    return vector_indexes_;
  }

  int Delete(int docid);

  std::map<std::string, RawVector *> RawVectors() { return raw_vectors_; }
  std::map<std::string, IndexModel *> IndexModels() { return vector_indexes_; }

  int MinIndexedNum();

  bitmap::BitmapManager *Bitmap() { return docids_bitmap_; };

  void Close();  // release all resource

 private:
  inline std::string IndexName(const std::string &field_name,
                               const std::string &index_type) {
    return field_name + "_" + index_type;
  }

 private:
  VectorStorageType default_store_type_;
  bitmap::BitmapManager *docids_bitmap_;
  bool table_created_;
  std::string root_path_;

  std::map<std::string, RawVector *> raw_vectors_;
  std::map<std::string, IndexModel *> vector_indexes_;
  std::vector<std::string> index_types_;
  std::vector<std::string> index_params_;
};

}  // namespace vearch
