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
#include "service303.h"
#include <gtest/gtest.h>
#include "MetricsRegistry.h"
#include <prometheus/registry.h>

using io::prometheus::client::MetricFamily;
using magma::service303::MetricsRegistry;
using prometheus::BuildCounter;
using prometheus::Registry;
using prometheus::detail::CounterBuilder;
using ::testing::Test;

// Tests the MetricsRegistry properly initializes and retrieves metrics
TEST_F(Test, TestMetricsRegistry)
{
  auto prometheus_registry = std::make_shared<Registry>();
  auto registry = MetricsRegistry<prometheus::Counter, CounterBuilder (&)()>(
    prometheus_registry, BuildCounter);
  EXPECT_EQ(registry.SizeFamilies(), 0);
  EXPECT_EQ(registry.SizeMetrics(), 0);

  // Create two new timeseries that will construct two families and metrics
  registry.Get("test", {});
  registry.Get("another", {{"key", "value"}});
  EXPECT_EQ(registry.SizeFamilies(), 2);
  EXPECT_EQ(registry.SizeMetrics(), 2);

  // This should retrieve the previously constructed family
  registry.Get("test", {});
  EXPECT_EQ(registry.SizeFamilies(), 2);
  EXPECT_EQ(registry.SizeMetrics(), 2);

  // Add new unique timeseries to an existing family
  registry.Get("test", {{"key", "value1"}});
  registry.Get("test", {{"key", "value2"}});
  EXPECT_EQ(registry.SizeFamilies(), 2);
  EXPECT_EQ(registry.SizeMetrics(), 4);
}

int main(int argc, char **argv)
{
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
