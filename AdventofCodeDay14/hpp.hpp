/*
** EPITECH PROJECT, 2023
** Adventofcode
** File description:
** hpp
*/

#ifndef HPP_HPP_
#define HPP_HPP_
#include <iostream>
#include <vector>
#include <fstream>

#define NORTH -1, 0
#define SOUTH 1, 0
#define EAST 0, 1
#define WEST 0, -1

class main_class {
    public:
        main_class(std::vector<std::string> input);
        ~main_class() = default;
        int get_result();
        void rotate_map(int x, int y);
        // void rotate_map_normal(int x, int y);
        // void rotate_map_not_normal(int x, int y);
        // void rotate_map_north(int x, int y);
        // void rotate_map_south(int x, int y);
        // void rotate_map_east(int x, int y);
        // void rotate_map_west(int x, int y);
    private:
        bool _stop;
        std::vector<std::vector<char>> _map;
        std::vector<std::vector<char>> _map_start;
        std::vector<std::vector<char>> _map_end;
};  

#endif //HPP_HPP_

