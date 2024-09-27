/**
 * Copyright 2019 The Gamma Authors.
 *
 * This source code is licensed under the Apache License, Version 2.0 license
 * found in the LICENSE file in the root directory of this source tree.
 */

#pragma once

#include <string>
#include <vector>

#include "rocksdb/options.h"
#include "rocksdb/table.h"
#include "vector/raw_vector.h"

namespace vearch {

class RocksDBRawVector : public RawVector {
 public:
  RocksDBRawVector(VectorMetaInfo *meta_info, const StoreParams &store_params,
                   bitmap::BitmapManager *docids_bitmap,
                   StorageManager *storage_mgr, int cf_id);
  ~RocksDBRawVector();
  /* RawVector */
  int InitStore(std::string &vec_name) override;
  int AddToStore(uint8_t *v, int len) override;
  int GetVectorHeader(int start, int n, ScopeVectors &vecs,
                      std::vector<int> &lens) override;
  int UpdateToStore(int vid, uint8_t *v, int len) override;

  size_t GetStoreMemUsage() override;

  int Gets(const std::vector<int64_t> &vids, ScopeVectors &vecs) const override;

  Status Dump(int start_vid, int end_vid) override { return Status::OK(); };
  int GetDiskVecNum(int &vec_num) override;
  Status Load(int vec_num) override;

 protected:
  int GetVector(long vid, const uint8_t *&vec, bool &deletable) const override;

 private:
  void ToRowKey(int vid, std::string &key) const;

 private:
  rocksdb::BlockBasedTableOptions table_options_;
  size_t block_cache_size_;
};
}  // namespace vearch
