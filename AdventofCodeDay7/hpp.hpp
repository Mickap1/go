/*
** EPITECH PROJECT, 2023
** camel_card
** File description:
** hpp
*/

#ifndef HPP_HPP
#define HPP_HPP

#include <iostream>
#include <vector>
#include <fstream>
#include <sstream>

enum hand_type {
    HIGH_CARD,
    ONE_PAIR,
    TWO_PAIRS,
    THREE_OF_A_KIND,
    FULL_HOUSE,
    FOUR_OF_A_KIND,
    FIVE_OF_A_KIND
};

enum card_type {
    J,
    TWO,
    THREE,
    FOUR,
    FIVE,
    SIX,
    SEVEN,
    EIGHT,
    NINE,
    TEN,
    Q,
    K,
    A
};

struct hand {
    std::vector<card_type> hand;
    int bid;
    hand_type type;

    // bool operator<(const struct hand& rhs) const{
    //     if (type == rhs.type) {
    //         for (int i = 0; i < 5; i++) {
    //             if (hand[i] < rhs.hand[i])
    //                 return true;
    //             else if (hand[i] > rhs.hand[i])
    //                 return false;
    //         }
    //     } else
    //         return type < rhs.type;
    // std::cout << "Error: hands are equal\n" << std::endl;
    // return true; //They are equal
    // }

    bool operator<(const struct hand& rhs) const {
    if (type != rhs.type) {
        return type < rhs.type;
    }

    for (int i = 0; i < 5; i++) {
        if (hand[i] != rhs.hand[i]) {
            return hand[i] < rhs.hand[i];
        }
    }

    std::cout << "Error: hands are equal\n" << std::endl;
    return false; // They are equal
    }
};

class camel_card {
    public:
        camel_card(std::string file_name);
        ~camel_card() = default;
        void get_file();
        void rank_hands();
        int get_result();
    private:
        std::vector<hand> hands;
        std::vector<hand> ranked_hands;
};

#endif //HPP_HPP
