/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the Apache License, Version 2.0  (the "License"); you may not use this file
 * except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */
#include <string.h>
#include <gtest/gtest.h>
#include "IMSIEncoder.h"

using ::testing::Test;
using ::testing::Values;
using namespace openflow;

namespace {

class IMSIEncoderTest : public ::testing::TestWithParam<std::string> {
};

/*
 * Test IMSI encoder within GTP application by encoding an IMSI to uint64_t and
 * back to see if the values match
 */
TEST_P(IMSIEncoderTest, TestCompactExpand)
{
  std::string imsi_test =
    IMSIEncoder::expand_imsi(IMSIEncoder::compact_imsi(GetParam()));
  ASSERT_STREQ(GetParam().c_str(), imsi_test.c_str());
}

INSTANTIATE_TEST_CASE_P(
  TestLeadingZeros,
  IMSIEncoderTest,
  Values("001010000000013", "011010000000013", "111010000000013"));

INSTANTIATE_TEST_CASE_P(
  TestDifferentLengths,
  IMSIEncoderTest,
  Values("001010000000013", "01010000000013", "28950000000013"));

int main(int argc, char **argv)
{
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}

} // namespace
