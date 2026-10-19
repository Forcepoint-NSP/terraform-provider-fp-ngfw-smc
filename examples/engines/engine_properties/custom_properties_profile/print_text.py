#!/usr/bin/env python3
"""Simple script that prints the text passed via --text."""

import argparse


def main():
    parser = argparse.ArgumentParser(description="Print the provided text.")
    parser.add_argument("--text", default="Hello World", help="The text to print.")
    args = parser.parse_args()
    print(args.text)


if __name__ == "__main__":
    main()
