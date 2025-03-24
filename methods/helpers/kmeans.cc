#pragma once

#include <vector>
#include "utils/rand.h"

void kmeans(const std::vector<float> *data, int n_centroids, int sample_size, int n, int d, int k, int rounds) {
    // Initialize random centroids
    std::vector<float> *centroids = new std::vector<float>(n_centroids*d);
    int seed = 1;
    for (int i = 0; i < k; i++) {
        int index = randInt(0, n-1, seed);
        for (int j = 0; j < d; j++) {
            (*centroids)[i*d+j] = (*data)[index*d+j];
        }
        seed++;
    }
    // extract sample
    std::vector<float> *sample = new std::vector<float>(sample_size*d);
    select_sample(data, sample_size, n, d, sample);

    // train centroids on the samplefor the given number of rounds
    std::vector<int> assignments(sample_size);
    for(int i = 0; i < rounds; i++) {
        assign_points(sample, centroids, sample_size, d, k, assignments);
        update_centroids(sample, centroids, sample_size, d, k, assignments);
    }
    // assign the rest of the data to the nearest centroid
    assign_points(data, centroids, n, d, k, assignments);
}

void select_sample(const std::vector<float> *data, int sample_size, int n, int d, std::vector<float> *sample) {
    int divisor = n/sample_size;
    for(int i = 0; i < sample_size; i++) {
        for(int j = 0; j < d; j++) {
            (*sample)[i*d+j] = (*data)[i*divisor*d+j];
        }
    }
}

void assign_points(const std::vector<float> *data, std::vector<float> *centroids, int n, int d, int k, std::vector<int> &assignments) {
    for(int i = 0; i < n; i++) {
        float min_distance = MAXFLOAT;
        int min_index = 0;
        for(int j = 0; j < k; j++) {
            float distance = 0;
            for (int l = 0; l < d; l++){
                distance += ((*data)[i*d+l] - (*centroids)[j*d+l]) * ((*data)[i*d+l] - (*centroids)[j*d+l]);
            }
            if(distance < min_distance) {
                min_distance = distance;
                min_index = j;
            }
        }
        assignments[i] = min_index;
    }
}

void update_centroids(const std::vector<float> *data, std::vector<float> *centroids, int n, int d, int k, std::vector<int> &assignments) {
    // TODO: implement update_centroids
}


