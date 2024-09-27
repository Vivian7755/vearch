/**
 * Copyright 2019 The Gamma Authors.
 *
 * This source code is licensed under the Apache License, Version 2.0 license
 * found in the LICENSE file in the root directory of this source tree.
 */

#pragma once

#include <string>

#include "raw_vector.h"

namespace vearch {

class MemoryRawVector : public RawVector {
 public:
  MemoryRawVector(VectorMetaInfo *meta_info, const StoreParams &store_params,
                  bitmap::BitmapManager *docids_bitmap,
                  StorageManager *storage_mgr, int cf_id);

  ~MemoryRawVector();

  int InitStore(std::string &vec_name) override;

  int AddToStore(uint8_t *v, int len) override;

  int GetVectorHeader(int start, int n, ScopeVectors &vecs,
                      std::vector<int> &lens) override;

  int UpdateToStore(int vid, uint8_t *v, int len) override;

  uint8_t *GetFromMem(long vid) const;

  int AddToMem(uint8_t *v, int len);

  Status Load(int vec_num) override;
  int GetDiskVecNum(int &vec_num) override;
  Status Dump(int start_vid, int end_vid) override { return Status::OK(); };

 protected:
  int GetVector(long vid, const uint8_t *&vec, bool &deleteable) const override;

 private:
  int ExtendSegments();

  uint8_t **segments_;
  int nsegments_;
  int segment_size_;
  uint8_t *current_segment_;
  int curr_idx_in_seg_;
};

}  // namespace vearch
