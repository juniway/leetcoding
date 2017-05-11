#!/usr/bin/env python


import os
import sys
import re


prefix = "    * ["


def walk_dir(id, dirname):
    i = 1
    fw = open("tmp.txt", 'w')
    for root, dirs, files in os.walk(dirname):
        # print root
        # print dirs
        for f in files:
            doti = f.find(".")
            num = f[:doti].lstrip("0")
            fa = f[doti:]
            words = re.findall('[A-Z0-9][a-z]*', fa)
            wds = ' '.join(map(str, words))
            # print f
            line = prefix + id + "." + repr(i) + " (" + num + ") " + wds + "](" + os.path.join(root, f) + ")\n"
            fw.write(line)
            i += 1
            # fp = os.path.join(root, f)
    fw.close()


def parse_argv(argv):
    if len(argv) < 3:
        print "arg len < 2"
        sys.exit()
    walk_dir(argv[2], argv[1])


if __name__ == "__main__":
    parse_argv(sys.argv)
