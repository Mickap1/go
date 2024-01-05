/*
** EPITECH PROJECT, 2023
** Adventofcode
** File description:
** cpp
*/

#include "hpp.hpp"
#include <cstdlib>

bool is_out_of_bound(std::vector<std::vector<char>> map, int x , int y)
{
    if (x < 0 || y < 0 || x >= map.size() || y >= map[0].size()) {
        // std::cout << "Out of bound" << std::endl;
        return true;
    }
    // std::cout << "Not out of bound" << std::endl;
    return false;
}

void display_map(std::vector<std::vector<char>> map)
{
    for (int i = 0; i < map.size(); i++) {
        for (int j = 0; j < map[i].size(); j++)
            std::cout << map[i][j];
        std::cout << std::endl;
    }
    std::cout << std::endl;
}

void main_class::rotate_map(int x, int y)
{
    bool has_moved = true;
    while (has_moved == true) {
        has_moved = false;
        for (int i = 0; i < this->_map.size(); i++) {
            for (int j = 0; j < this->_map[i].size(); j++) {
                if (this->_map[i][j] == 'O') {
                    int move = 1;
                    while (is_out_of_bound(this->_map, i + (move * x), j + (move * y)) == false && this->_map[i + (move * x)][j + (move * y)] == '.') {
                        move++;
                    }
                    move--;
                    this->_map[i][j] = '.';
                    this->_map[i + (move * x)][j + (move * y)] = 'O';
                    if (move >= 1)
                        has_moved = true;
                }
            }
        }
    }
}

main_class::main_class(std::vector<std::string> input)
{
    this->_stop = false;
    for (int i = 0; i < input.size(); i++) {
        std::vector<char> line;
        for (int j = 0; j < input[i].size(); j++) {
            line.push_back(input[i][j]);
        }
        this->_map.push_back(line);
    }
    for (int i = 0; i < 1000000000 && this->_stop == false; i++) {
        this->_map_start = this->_map;
        rotate_map(NORTH);
        rotate_map(WEST);
        rotate_map(SOUTH);
        rotate_map(EAST);
        this->_map_end = this->_map;
        if (this->_map_start == this->_map_end) {
            std::cout << "Found a loop at " << i << std::endl;
            this->_stop = true;
        }
    }
}

std::vector<std::string> open_file(std::string filename)
{
    std::vector<std::string> lines;
    std::ifstream file(filename);

    std::string line;
    while (std::getline(file, line)) {
        lines.push_back(line);
    }

    file.close();

    return (lines);
}

int main_class::get_result()
{
    int result = 0;
    int nb_of_lines = this->_map.size();
    int nb_of_rock = 0;
    for (int i = 0; i < nb_of_lines; i++) {
        nb_of_rock = 0;
        for (int j = 0; j < this->_map[i].size(); j++) {
            if (this->_map[i][j] == 'O')
                nb_of_rock++;
        }
        result += nb_of_rock * (nb_of_lines - i);
    }
    return (result);
}

int main(int ac, char **av)
{
    if (ac != 2)
        return 84;
    int result = 0;
    std::vector<std::string> lines = open_file(av[1]);
    main_class main(lines);
    result = main.get_result();
    std::cout << "Result: " << result << std::endl;
    return 0;
}
