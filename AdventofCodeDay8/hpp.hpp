/*
** EPITECH PROJECT, 2023
** Adventofcode
** File description:
** hpp
*/

#ifndef HPP_HPP
#define HPP_HPP

#include <iostream>
#include <fstream>
#include <string>
#include <vector>

class box {
    public:
        box(std::string name, std::string right, std::string left);
        std::string name;
        std::string right;
        std::string left;
};

class all_box {
    public:
        all_box(std::string path);
        int get_result(void); 
    private:
        int find_result(std::string c);
        std::vector<box> box_list;
        std::vector<char> instructions;
        // int find_resul(std::string c);
        int result;
};

#endif
