/*
** EPITECH PROJECT, 2023
** camel_card
** File description:
** main
*/

#include "hpp.hpp"
#include <algorithm>
#include <string>

void print_hand(const hand& hand) {
    // std::cout << "Bid: " << hand.bid << std::endl;
    std::cout << "Hand: ";
    for (card_type c : hand.hand) {
        std::cout << c << " ";
    }
    std::cout << std::endl;
    // std::cout << "Type: " << hand.type << std::endl;
}

bool same_card(card_type a, card_type b) {
    if (a == J)
        return true;
    if (b == J)
        return true;
    return a == b;
}

void get_all_hands(std::vector<std::vector<card_type>> *all_hands, std::vector<card_type> base_hand, std::vector<card_type> diff_joker, std::vector<int> index_joker) {
    for (int i = 0; i < diff_joker.size(); i++) {
        std::vector<card_type> hand_copy = base_hand;
        hand_copy[index_joker[0]] = diff_joker[i];
        index_joker.erase(index_joker.begin());
        get_all_hands(all_hands, hand_copy, diff_joker, index_joker);
    }
}

hand_type get_type(std::vector<card_type> hand) {
    std::vector<card_type> hand_copy = hand;

    std::vector<card_type> diff_joker;
    bool is_joker = false;
    std::vector<int> index_joker;

    for (int i = 0; i < 5; i++) {
        if (hand_copy[i] == J) {
            is_joker = true;
            index_joker.push_back(i);
        } else {
            diff_joker.push_back(hand_copy[i]);
        } 
    }
    if (is_joker && diff_joker.size() == 0)  {
        return FIVE_OF_A_KIND;
    }
    if (is_joker) {
        if (diff_joker.size() <= 1) {
            return FIVE_OF_A_KIND;
        }
        std::vector<card_type> hand_copy2 = hand_copy;
        hand_type best_type = HIGH_CARD;
        std::vector<std::vector<card_type>> all_hands;
        get_all_hands(&all_hands, hand_copy, diff_joker, index_joker);
        return best_type;
    }

    std::sort(hand_copy.begin(), hand_copy.end());


    if (same_card(hand_copy[0], hand_copy[1]) && same_card(hand_copy[1], hand_copy[2]) && same_card(hand_copy[2], hand_copy[3]) && same_card(hand_copy[3], hand_copy[4]))
        return FIVE_OF_A_KIND;

    if (same_card(hand_copy[0], hand_copy[1]) && same_card(hand_copy[1], hand_copy[2]) && same_card(hand_copy[2], hand_copy[3]))
        return FOUR_OF_A_KIND;

    if (same_card(hand_copy[1], hand_copy[2]) && same_card(hand_copy[2], hand_copy[3]) && same_card(hand_copy[3], hand_copy[4]))
        return FOUR_OF_A_KIND;

    if (same_card(hand_copy[0], hand_copy[1]) && same_card(hand_copy[1], hand_copy[2]) && same_card(hand_copy[3], hand_copy[4]))
        return FULL_HOUSE;

    if (same_card(hand_copy[0], hand_copy[1]) && same_card(hand_copy[2], hand_copy[3]) && same_card(hand_copy[3], hand_copy[4]))
        return FULL_HOUSE;

    if (same_card(hand_copy[0], hand_copy[1]) && same_card(hand_copy[1], hand_copy[2]))
        return THREE_OF_A_KIND;

    if (same_card(hand_copy[1], hand_copy[2]) && same_card(hand_copy[2], hand_copy[3]))
        return THREE_OF_A_KIND;

    if (same_card(hand_copy[2], hand_copy[3]) && same_card(hand_copy[3], hand_copy[4]))
        return THREE_OF_A_KIND;

    if (same_card(hand_copy[0], hand_copy[1]) && same_card(hand_copy[2], hand_copy[3]))
        return TWO_PAIRS;

    if (same_card(hand_copy[0], hand_copy[1]) && same_card(hand_copy[3], hand_copy[4]))
        return TWO_PAIRS;

    if (same_card(hand_copy[1], hand_copy[2]) && same_card(hand_copy[3], hand_copy[4]))
        return TWO_PAIRS;

    if (same_card(hand_copy[0], hand_copy[1]))
        return ONE_PAIR;

    if (same_card(hand_copy[1], hand_copy[2]))
        return ONE_PAIR;

    if (same_card(hand_copy[2], hand_copy[3]))
        return ONE_PAIR;

    if (same_card(hand_copy[3], hand_copy[4]))
        return ONE_PAIR;

    return HIGH_CARD;
}

card_type transformCharToEnum(char c) {
    switch (c) {
        case '2':
            return TWO;
        case '3':
            return THREE;
        case '4':
            return FOUR;
        case '5':
            return FIVE;
        case '6':
            return SIX;
        case '7':
            return SEVEN;
        case '8':
            return EIGHT;
        case '9':
            return NINE;
        case 'T':
            return TEN;
        case 'J':
            return J;
        case 'Q':
            return Q;
        case 'K':
            return K;
        case 'A':
            return A;
        default:
            std::cout << "Invalid card" << std::endl;
            throw std::invalid_argument("Invalid card");
    }
}

hand transformStringToStruct(const std::string& input) {
    hand result;

    for (int i = 0; i < 5; i++) {
        result.hand.push_back(transformCharToEnum(input[i]));
    }
    std::stringstream ss(input.substr(5));
    ss >> result.bid;
    result.type = get_type(result.hand);

    return result;
}

camel_card::camel_card(std::string file_name)
{
    std::ifstream file(file_name);
    std::string line;
    std::vector<std::string> hand;
    std::vector<char> hand_cards;
    int bid;

    if (!file.is_open())
        throw std::invalid_argument("File not found");
    while (std::getline(file, line)) {
        this->hands.push_back(transformStringToStruct(line));
    }
    this->ranked_hands = this->hands;
    std::sort(this->ranked_hands.begin(), this->ranked_hands.end());

    file.close();
}

int camel_card::get_result()
{
    long total = 0;
    for (int i = 0; i < this->ranked_hands.size(); i++) {
        total += this->ranked_hands[i].bid * (i + 1);
        // print_hand(this->ranked_hands[i]);
        // std::cout << total << std::endl;
        // if (this->ranked_hands[i].bid == 553)
        //     std::cout << this->ranked_hands[i].type << std::endl;
        // if (this->ranked_hands[i].bid == 394)
        //     std::cout << this->ranked_hands[i].type << std::endl;
        // std::cout << this->ranked_hands[i].bid << std::endl;
    }
    return total;
}

int main(int ac, char **av)
{
    // std::string path = "test2.txt";
    camel_card class_name(av[1]);
    std::cout << class_name.get_result() << std::endl;
    // std::string path2 = "test.txt";
    // camel_card class_name2(av[2]);
    // std::cout << class_name2.get_result() << std::endl;
    return (0);
}
