/*
** EPITECH PROJECT, 2023
** camel_card
** File description:
** main
*/

#include "hpp.hpp"
// #include <algorithm>
#include <string>



std::vector<std::string> ends_with(std::vector<box> box_list, std::string end)
{
    std::vector<std::string> names;
    for (int i = 0; i < box_list.size(); i++) {
        if (box_list[i].name[2] == end[0])
            names.push_back(box_list[i].name);
    }
    return (names);
}

bool do_they_all_end_with(std::vector<std::string> names, std::string end, int current_round)
{
    int size = names.size();
    int fnd_count = 0;
    for (int i = 0; i < names.size(); i++) {
        if (names[i][2] == end[0])
            fnd_count++;
    }
    // if (fnd_count > 0)
    //     std::cout << "fnd_count: " << fnd_count << "out of " << size <<" round: " << current_round <<std::endl;
    if (fnd_count == size)
        return (true);
    return (false);
}

box::box(std::string name, std::string right, std::string left)
{
    this->name = name;
    this->right = right;
    this->left = left;
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

int all_box::find_result(std::string c)
{
    std::string current_name = c;
    std::string end_name = "ZZZ";
    int _result = 0;
    int index_current_instruction = 0;
    while (do_they_all_end_with({current_name}, end_name, 0) == false) {
        if (index_current_instruction == instructions.size())
            index_current_instruction = 0;
        if (this->instructions[index_current_instruction] == 'R') {
            for (int i = 0; i < box_list.size(); i++) {
                if (box_list[i].name == current_name) {
                    // std::cout << "RIGHT of " << current_name << " is " << box_list[i].right << std::endl;
                    current_name = box_list[i].right;
                    break;
                }
            }
        } else if (this->instructions[index_current_instruction] == 'L') {
            for (int i = 0; i < box_list.size(); i++) {
                if (box_list[i].name == current_name) {
                    // std::cout << "LEFT of " << current_name << " is " << box_list[i].left << std::endl;
                    current_name = box_list[i].left;
                    break;
                }
            }
        } else {
            std::cout << "Error: instruction not found" << std::endl;
            exit(84);
        }
        index_current_instruction++;
        _result++;
    } 
    return (_result);
}

// int all_box::find_result(std::string name)
// {
//     std::string current_name = name;
//     int index_current_instruction = 0;
//     while (name[2] != "Z") {
//         if (index_current_instruction == instructions.size())
//             index_current_instruction = 0;
//         if (this->instructions[index_current_instruction] == 'R') {
//             for (int i = 0; i < box_list.size(); i++) {
//                 if (box_list[i].name == current_name) {
//                     // std::cout << "RIGHT of " << current_name << " is " << box_list[i].right << std::endl;
//                     current_name = box_list[i].right;
//                     break;
//                 }
//             }
//         } else if (this->instructions[index_current_instruction] == 'L') {
//             for (int i = 0; i < box_list.size(); i++) {
//                 if (box_list[i].name == current_name) {
//                     // std::cout << "LEFT of " << current_name << " is " << box_list[i].left << std::endl;
//                     current_name = box_list[i].left;
//                     break;
//                 }
//             }
//         } else {
//             std::cout << "Error: instruction not found" << std::endl;
//             exit(84);
//         }
//         }
//         index_current_instruction++;
//         this->result++;
//     } 
//     // for (int i = 0; i < current_names.size(); i++)
//     //     std::cout << current_names[i] << std::endl;
// }

// void all_box::find_result(void)
// {
//     std::vector<std::string> current_names = ends_with(this->box_list, "A");
//     std::vector<int> first_find_index;
//     std::vector<int> next_finds
//     ;
// }

all_box::all_box(std::string path)
{
    this->result = 0;
    std::vector<std::string> lines = open_file(path);
    for (int i = 0; i < lines[0].size(); i++)
        instructions.push_back(lines[0][i]);
    for (int i = 0; i < this->instructions.size(); i++)
        std::cout << instructions[i];
    std::cout << std::endl;
    for (int i = 2; i < lines.size(); i++) {
        std::string name = "";
        std::string right = "";
        std::string left = "";
        for (int j = 0; j < 3; j++)
            name += lines[i][j];
        for (int j = 7; j < 10; j++)
            right += lines[i][j];
        for (int j = 12; j < 15; j++)
            left += lines[i][j];
        box_list.push_back(box(name, left, right));
    }
    std::vector<std::string> current_names = ends_with(this->box_list, "A");
    std::vector<int> results;
    for (int i = 0; i < current_names.size(); i++) {
        results.push_back(this->find_result(current_names[i]));
        std::cout << results[i] << " * ";
    }
    std::cout << std::endl;
    this->result = results[0];
    for (int i = 1; i < results.size(); i++) {
        this->result *= results[i];
    }
    std::cout << "Result shoyld be :"   << this->result << std::endl;
}

int all_box::get_result(void)
{
    return (this->result);
}

int main(int ac, char **av)
{
    all_box myboxs(av[1]);
    std::cout << myboxs.get_result() << std::endl;
    return (0);
}

//No longer working x')
