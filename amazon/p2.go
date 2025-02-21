package main

// Complete the 'findMaximumSustainableClusterSize' function below.
//
// The function is expected to return an INTEGER.
// The function accepts the following parameters:
//  1. INTEGER_ARRAY processingPower
//  2. INTEGER_ARRAY bootingPower
//  3. LONG_INTEGER powerMax

import (
	"github.com/gammazero/deque"
)

// findMaximumSustainableClusterSizeNaive finds maximum sustainable cluster
// size using a sliding window.
func findMaximumSustainableClusterSizeNaive(
	processingPower []int32, bootingPower []int32, powerMax int64) int32 {
	// This problem can be approached as a sliding window problem given the
	// constraint:
	//
	//   >Clusters can only be formed of processors located adjacent to each
	//    other.
	//
	// Therefore, starting at index zero, the window is adjusted such that the
	// maximum cluster is formed when the following criteria is met:
	//
	//   net power consumption = maximum booting power + (sum of processing power) * k
	//
	// where net power consumption ≤ max power.
	maxClusterSize := 0
	numProcessors := len(processingPower)
	// Calculate the largest sustainable cluster starting with processor i
	for i := 0; i < numProcessors; i++ {
		clusterSize := 1
		sumProcessingPower := int64(processingPower[i])
		maxBootingPower := int64(bootingPower[i])

		// Calculate the inital net power consumption
		power := maxBootingPower + sumProcessingPower*int64(clusterSize)
		if power <= powerMax && clusterSize > maxClusterSize {
			maxClusterSize = clusterSize
		}

		// Attempt to add processor j to the cluster.
		for j := i + 1; j < numProcessors; j++ {
			// Recalculate maximum booting power
			if int64(bootingPower[j]) > maxBootingPower {
				maxBootingPower = int64(bootingPower[j])
			}
			// Recalculate sum of processing power
			sumProcessingPower += int64(processingPower[j])
			clusterSize++
			power = maxBootingPower + sumProcessingPower*int64(clusterSize)
			// If the cluster is sustainable (net power consumption does not exceed
			// powerMax) and is greater than the current maximum cluster size, update
			// maxClusterSize
			if power <= powerMax && clusterSize > maxClusterSize {
				maxClusterSize = clusterSize
			}
		}
	}
	return int32(maxClusterSize)
}

// findMaximumSustainableClusterSizeOptimal finds maximum sustainable cluster
// size using a monotonic queue.
func findMaximumSustainableClusterSizeOptimal(
	processingPower []int32, bootingPower []int32, powerMax int64) int32 {
	// A monotonic queue (or monotone priority queue) is maintained with the
	// maximum booting power of the current group of processors.
	//
	// The monotonic queue is implemented using a deque in order to efficiently
	// add and remove items at either end with O(1) performance.
	//
	// Pointers to the start and end of the sliding window are maintained in
	// order to preserve the following criteria:
	//
	//   net power consumption = maximum booting power + (sum of processing power) * k
	//
	// where net power consumption ≤ max power.
	//
	// NOTE: In order to preserve the bounds of the sliding window, the deque
	// holds the indices of bootingPower, not its values.
	var mq deque.Deque[int]
	maxClusterSize, clusterSize, start, end := 0, 0, 0, 0
	var sumProcessingPower int32 = 0
	for end < len(processingPower) {
		// Remove elements from the back of the queue until the queue is empty or
		// the element at the back of the queue is greater than the current
		// element, then append the current element to the back of the queue.
		for mq.Len() > 0 && bootingPower[mq.Back()] < bootingPower[end] {
			mq.PopBack()
		}
		mq.PushBack(end)
		maxBootingPower := bootingPower[mq.Front()]
		sumProcessingPower += processingPower[end]
		power := maxBootingPower + sumProcessingPower*int32(clusterSize)
		clusterSize++
		if int64(power) <= powerMax {
			end++
		} else {
			sumProcessingPower -= processingPower[start]
			clusterSize--
			start++
			end++
		}
		if clusterSize > maxClusterSize {
			maxClusterSize = clusterSize
		}
		// Preserve the bounds of the sliding window by checking that the element
		// at the front of the queue is still within the window. If not, remove it.
		if start > mq.Front() {
			mq.PopFront()
		}
	}
	return int32(maxClusterSize)
}
