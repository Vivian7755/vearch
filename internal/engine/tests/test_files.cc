/**
 * Copyright 2019 The Gamma Authors.
 *
 * This source code is licensed under the Apache License, Version 2.0 license
 * found in the LICENSE file in the root directory of this source tree.
 */

#include "test.h"

/**
 * To run this demo, please download the ANN_SIFT10K dataset from
 *
 *   ftp://ftp.irisa.fr/local/texmex/corpus/siftsmall.tar.gz
 *
 * and unzip it.
 **/

namespace test {

class GammaTest : public ::testing::Test {
 public:
  static int Init(int argc, char *argv[]) {
    GammaTest::my_argc = argc;
    GammaTest::my_argv = argv;
    return 0;
  }

 protected:
  GammaTest() {}

  ~GammaTest() override {}

  void SetUp() override {}

  void TearDown() override {}

  void *engine;

  static int my_argc;
  static char **my_argv;
};

int GammaTest::my_argc = 0;
char **GammaTest::my_argv = nullptr;

TEST_F(GammaTest, IVFPQ) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFPQ_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFPQ_MEMORYONLY) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFPQParam;
  opt.store_type = "MemoryOnly";
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFPQ_MEMORYONLY_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFPQParam;
  opt.store_type = "MemoryOnly";
  ASSERT_EQ(TestIndexes(opt), 0);
  
  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFPQ_ROCKSDB) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFPQParam;
  opt.store_type = "RocksDB";
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFPQ_ROCKSDB_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFPQParam;
  opt.store_type = "RocksDB";
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFPQ_HNSW) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFHNSWPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFPQ_HNSW_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFHNSWPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFPQ_HNSW_OPQ) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFHNSWOPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFPQ_HNSW_OPQ_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFPQ";
  opt.index_params = kIVFHNSWOPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFFLAT) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFFLAT";
  opt.index_params = kIVFPQParam;
  opt.store_type = "RocksDB";
  ASSERT_EQ(TestIndexes(opt), 0);
  
  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, IVFFLAT_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "IVFFLAT";
  opt.index_params = kIVFPQParam;
  opt.store_type = "RocksDB";
  ASSERT_EQ(TestIndexes(opt), 0);
  
  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, FLAT) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "FLAT";
  opt.index_params = kIVFPQParam;
  opt.store_type = "MemoryOnly";
  opt.add_doc_num = 100000;
  ASSERT_EQ(TestIndexes(opt), 0);
  
  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, FLAT_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "FLAT";
  opt.index_params = kIVFPQParam;
  opt.store_type = "MemoryOnly";
  opt.add_doc_num = 100000;
  ASSERT_EQ(TestIndexes(opt), 0);
  
  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, HNSW) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "HNSW";
  opt.index_params = kHNSWParam;
  opt.store_type = "MemoryOnly";
  opt.add_doc_num = 20000;
  ASSERT_EQ(TestIndexes(opt), 0);
  
  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, HNSW_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "HNSW";
  opt.index_params = kHNSWParam;
  opt.store_type = "MemoryOnly";
  opt.add_doc_num = 20000;
  ASSERT_EQ(TestIndexes(opt), 0);
  
  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}
/*
TEST_F(GammaTest, BINARYIVF) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "BINARYIVF";
  opt.index_params = kIVFBINARYParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, BINARYIVF_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "BINARYIVF";
  opt.index_params = kIVFBINARYParam;
  ASSERT_EQ(TestIndexes(opt), 0);
  
  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}
*/
#ifdef USE_SCANN
TEST_F(GammaTest, SCANN_ROCKSDB) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "VEARCH";
  opt.index_params = kSCANNParam;
  opt.store_type = "RocksDB";
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, SCANN_THREADPOOL_ROCKSDB) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "VEARCH";
  opt.index_params = kSCANNWithThreadPoolParam;
  opt.store_type = "RocksDB";
  ASSERT_EQ(TestIndexes(opt), 0);
}
#endif

#ifdef BUILD_WITH_GPU
TEST_F(GammaTest, GPU) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "GPU";
  opt.index_params = kIVFPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, GPU_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "GPU";
  opt.index_params = kIVFPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, GPU_MEMORYONLY) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "GPU";
  opt.store_type = "MemoryOnly";
  opt.index_params = kIVFPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, GPU_MEMORYONLY_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "GPU";
  opt.store_type = "MemoryOnly";
  opt.index_params = kIVFPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, GPU_ROCKSDB) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "GPU";
  opt.store_type = "RocksDB";
  opt.index_params = kIVFPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}

TEST_F(GammaTest, GPU_ROCKSDB_BATCH) {
  struct Options opt;
  opt.set_file(my_argv, my_argc);
  opt.index_type = "GPU";
  opt.store_type = "RocksDB";
  opt.index_params = kIVFPQParam;
  ASSERT_EQ(TestIndexes(opt), 0);

  opt.b_load = true;
  ASSERT_EQ(TestIndexes(opt), 0);
}
#endif

}  // namespace test

int main(int argc, char **argv) {
  setvbuf(stdout, (char *)NULL, _IONBF, 0);
  ::testing::InitGoogleTest(&argc, argv);
  if (argc != 3 && argc != 4) {
    std::cout << "Usage: [Program] [profile_file] [vectors_file]\n";
    std::cout << "Usage: [Program] [profile_file] [vectors_file] [raw_data_type]\n";
    return 1;
  }
  ::testing::GTEST_FLAG(output) = "xml";
  test::GammaTest::Init(argc, argv);
  return RUN_ALL_TESTS();
}
