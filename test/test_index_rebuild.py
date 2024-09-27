#
# Copyright 2019 The Vearch Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
# implied. See the License for the specific language governing
# permissions and limitations under the License.

# -*- coding: UTF-8 -*-

import requests
import json
import pytest
from utils.vearch_utils import *
from utils.data_utils import *

__description__ = """ test case for index rebuild """


sift10k = DatasetSift10K()
xb = sift10k.get_database()
xq = sift10k.get_queries()
gt = sift10k.get_groundtruth()


class TestIndexRebuild:
    def setup_class(self):
        pass

    def test_prepare_db(self):
        logger.info(create_db(router_url, db_name))

    @pytest.mark.parametrize(
        ["training_threshold", "index_type"],
        [[1, "FLAT"], [9999, "IVFPQ"], [9999, "IVFFLAT"], [1, "HNSW"]],
    )
    def test_space_create(self, training_threshold, index_type):
        embedding_size = xb.shape[1]
        batch_size = 100
        total = xb.shape[0]
        total_batch = int(total / batch_size)
        with_id = True
        full_field = True
        logger.info(
            "dataset num: %d, total_batch: %d, dimension: %d" % (
                total, total_batch, embedding_size)
        )

        space_config = {
            "name": space_name,
            "partition_num": 1,
            "replica_num": 1,
            "fields": [
                {
                    "name": "field_int",
                    "type": "integer",
                },
                {
                    "name": "field_long",
                    "type": "long",
                },
                {
                    "name": "field_float",
                    "type": "float",
                },
                {
                    "name": "field_double",
                    "type": "double",
                },
                {
                    "name": "field_string",
                    "type": "string",
                    "index": {
                        "name": "field_string",
                        "type": "SCALAR",
                    },
                },
                {
                    "name": "field_vector",
                    "type": "vector",
                    "index": {
                        "name": "gamma",
                        "type": index_type,
                        "params": {
                            "metric_type": "InnerProduct",
                            "ncentroids": 256,
                            "nsubvector": 32,
                            "nlinks": 32,
                            "efConstruction": 40,
                            "training_threshold": training_threshold
                        },
                    },
                    "dimension": embedding_size,
                    # "format": "normalization"
                },
            ]
        }

        response = create_space(router_url, db_name, space_config)
        logger.info(response)
        add(total_batch, batch_size, xb, with_id, full_field)

        if index_type != "FLAT":
            waiting_index_finish(total)

        response = index_rebuild(router_url, db_name, space_name)
        logger.info(response)

        if index_type != "FLAT":
            waiting_index_finish(total)

        drop_space(router_url, db_name, space_name)

    def test_destroy_db(self):
        drop_db(router_url, db_name)
