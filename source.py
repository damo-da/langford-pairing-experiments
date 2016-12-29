#!/usr/bin/env python 
from copy import deepcopy
from time import sleep
import sys
import numpy as np

import datetime
from timer_helpers import TimeoutError, timeout

def solve(number, cur_list=None, exit_on_find=False):
    if cur_list is None:
        cur_list = [0] * (2*number)
    else:
        cur_list = deepcopy(cur_list)

    output = []

    empty_indices = [i for i,x in enumerate(cur_list) if x == 0]

    if not empty_indices:
        return cur_list
    else:
        for index in empty_indices:
            #print("Loop for {} at index {}".format(number, index))
            next_index = index + number + 1

            if next_index < len(cur_list) and cur_list[next_index] == 0:

                new_list = deepcopy(cur_list)
                new_list[index] = number
                new_list[next_index] = number

                #print("Filled {} on {} and {}".format(number, index, next_index))
                #print("Now, the list is {}".format(new_list))

                if number == 1:
                    output.append(new_list)
                    if exit_on_find:
                        return output
                else:
                    results = solve(number-1, new_list, exit_on_find=exit_on_find)
                    output += results

                    if len(results) > 0 and exit_on_find:
                        return output

            else:
                continue

    
    return output


def validate(result):
    assert(len(result)%2 == 0) #even number of items
    items_count = int(len(result)/2)

    assert ( 0 not in result)

    for i in range(1, items_count + 1):
        first_index = result.index(i)
        next_index = result.index(i, first_index+1)

        assert(next_index - first_index == i + 1)

    return True


def display(result):
    arr = [str(x) for x in result]

    print("Found {{{}}}".format(", ".join(arr)))
    sys.stdout.flush()


def for_single(num, exit_on_find=False):
    results = solve(num, exit_on_find=exit_on_find)

    print("#----set of {} has {} results".format(num, len(results)))
    sys.stdout.flush()

    for result in results:
        validate(result)
        display(result)


def for_range(start, end, exit_on_find=False):
    for i in range(start, end +1):
        if i % 4 not in (0,3):
            continue
        print("Looping for {}".format(i))
        startTime = datetime.datetime.now()
        
        for_single(i, exit_on_find=exit_on_find)
        endTime = datetime.datetime.now()

        print("Took {} time. ".format(endTime - startTime))
        #sleep(0.3)


def brutal_attack_like_vandal_savage(number):
    """THIS IS SPARTAAAAAA!"""

    @timeout(2)
    def hululu_solve(number, m_list, exit_on_find=True):
        """hululu hookah."""
        return solve(number, m_list, exit_on_find=exit_on_find)

    print("Solving for {}".format(number))
    for index in range(0, number-1):
        print("Working at index {}".format(index))

        m_list = [0] * (number * 2)

        m_list[index] = number
        m_list[index + number + 1] = number

        try:
            result = hululu_solve(number - 1, m_list)
            validate(result[0])
            print(result[0])
            return True
        except TimeoutError:
            print("TIMEOUT!!!@!!");

if __name__ == "__main__":
    #for_range(30,36, exit_on_find=True)
    brutal_attack_like_vandal_savage(47)

    #for_single(3, False)

