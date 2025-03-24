#pragma once
#include <random>

int randInt(int min, int max,int seed) {
    std::mt19937 gen(seed);
    std::uniform_int_distribution<> dis(min, max);
    return dis(gen);
        
}

