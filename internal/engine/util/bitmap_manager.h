/**
 * Copyright 2019 The Gamma Authors.
 *
 * This source code is licensed under the Apache License, Version 2.0 license
 * found in the LICENSE file in the root directory of this source tree.
 */

#pragma once

#include <string>

#include "rocksdb/db.h"
#include "rocksdb/options.h"
#include "rocksdb/table.h"
#include "util/status.h"

namespace bitmap {

class BitmapManager {
 public:
  BitmapManager();
  virtual ~BitmapManager();

  virtual int Init(uint32_t bit_size, const std::string &fpath = "",
                   std::shared_ptr<char[]> bitmap = nullptr);

  virtual int SetDumpFilePath(const std::string &fpath);

  virtual int Dump(uint32_t begin_bit_id = 0, uint32_t bit_len = 0);

  virtual int Load(uint32_t bit_len = 0);

  virtual uint32_t FileBytesSize();

  bool IsLoad() { return is_load_; }

  virtual int Set(uint32_t bit_id);

  virtual int Unset(uint32_t bit_id);

  virtual bool Test(uint32_t bit_id);

  virtual uint32_t BitSize() { return size_; }

  std::shared_ptr<char[]> Bitmap() { return bitmap_; }

  virtual uint32_t BytesSize() { return (size_ >> 3) + 1; }

  virtual int SetMaxID(uint32_t bit_id);

  std::shared_ptr<char[]> bitmap_;
  uint32_t size_;
  int fd_;
  std::string fpath_;
  bool is_load_;
};

constexpr uint32_t kBitmapSegmentBits = 1024 * 8;
constexpr uint32_t kBitmapSegmentBytes = 1024;
constexpr uint32_t kBitmapCacheSize = 10 * 1024 * 1024;
const std::string kBitmapSizeKey = "bitmap_size";

class RocksdbBitmapManager : public BitmapManager {
 public:
  RocksdbBitmapManager();
  virtual ~RocksdbBitmapManager();

  virtual int Init(uint32_t bit_size, const std::string &fpath = "",
                   std::shared_ptr<char[]> bitmap = nullptr);

  virtual int SetDumpFilePath(const std::string &fpath);

  virtual int Dump(uint32_t begin_bit_id = 0, uint32_t bit_len = 0);

  virtual int Load(uint32_t bit_len = 0);

  virtual uint32_t FileBytesSize();

  virtual int Set(uint32_t bit_id);

  virtual int Unset(uint32_t bit_id);

  virtual bool Test(uint32_t bit_id);

  virtual int SetMaxID(uint32_t bit_id);

  virtual void ToRowKey(uint32_t bit_id, std::string &key);

  rocksdb::DB *db_;
  bool should_load_;
};

}  // namespace bitmap
