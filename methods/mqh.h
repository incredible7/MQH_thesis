#pragma once

#include "helpers/kmeans.h"
#include "helpers/pq.h"


template<class DType>
class MQH {
public:
    int n_; // number of points
    int d_; // dimension of points
    const DType *data_; // pointerdata points


    // -------------------------------------------------------------------------
    MQH(const DType *data, int n, int d); // constructor
    // -------------------------------------------------------------------------
    ~MQH(); // destructor
    // -------------------------------------------------------------------------
    void build();
    // -------------------------------------------------------------------------
    int nns(const DType *query, int k); // nearest neighbor search
};

// -----------------------------------------------------------------------------

template<class DType>
MQH<DType>::MQH( // constructor
    const DType *data,
    int n,
    int d
) : data_(data), n_(n), d_(d) {
    
    build(); // build the datastructures
}
// -----------------------------------------------------------------------------

template <class DType>
MQH<DType>::~MQH() {
    // TODO: implement destructor
}


// -----------------------------------------------------------------------------

template <class DType>
int MQH<DType>::nns(const DType *query, int k) {
    // TODO: implement nearest neighbor search
}

// -----------------------------------------------------------------------------

void MQH<DType>::build() {
    build_quantization_tables();
    build_hash_tables();
}

// -----------------------------------------------------------------------------

void MQH<DType>::build_quantization_tables() {
    
    coarse_quantization();
    // TODO: implement build_quantization_tables
}

// -----------------------------------------------------------------------------

void MQH<DType>::build_hash_tables() {
    // TODO: implement build_hash_tables
}