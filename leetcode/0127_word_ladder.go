package main

import (
	"fmt"
)

/*
Word Ladder - LeetCode Problem 127 (Medium)

Problem:
Given two words, beginWord and endWord, and a dictionary wordList,
return the number of words in the shortest transformation sequence from
beginWord to endWord, or 0 if no such sequence exists.

You must change exactly one letter in each step, and each transformed word must
exist in the word list. Note that beginWord does not need to be in wordList.

Example:
Input: beginWord = "hit", endWord = "cog",
       wordList = ["hot","dot","dog","lot","log","cog"]
Output: 5
Explanation: "hit" -> "hot" -> "dot" -> "dog" -> "cog"

Approach: BFS (Breadth-First Search)
- Use BFS to find the shortest path from beginWord to endWord
- For each word, generate all possible one-letter transformations
- Check if the transformation exists in the word list
- Keep track of visited words to avoid cycles
- Return the path length when endWord is found

Time Complexity: O(N * L^2 * 26) where N is the number of words, L is the length of words
Space Complexity: O(N * L) for the queue and visited set
*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// Check if endWord exists in the list
	wordSet := make(map[string]bool)
	for _, word := range wordList {
		wordSet[word] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	// BFS setup
	type Node struct {
		word  string
		level int
	}

	queue := []Node{{beginWord, 1}}
	visited := make(map[string]bool)
	visited[beginWord] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.word == endWord {
			return current.level
		}

		// Generate all possible one-letter transformations
		for i := 0; i < len(current.word); i++ {
			for c := 'a'; c <= 'z'; c++ {
				if rune(current.word[i]) != c {
					// Create new word with character substitution
					bytes := []byte(current.word)
					bytes[i] = byte(c)
					newWord := string(bytes)

					// Check if the new word is valid and not visited
					if wordSet[newWord] && !visited[newWord] {
						visited[newWord] = true
						queue = append(queue, Node{newWord, current.level + 1})
					}
				}
			}
		}
	}

	return 0
}

func main() {
	// Test cases
	tests := []struct {
		beginWord string
		endWord   string
		wordList  []string
		expected  int
	}{
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			expected:  5,
		},
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			expected:  0,
		},
	}

	for i, test := range tests {
		result := ladderLength(test.beginWord, test.endWord, test.wordList)
		if result == test.expected {
			fmt.Printf("Test %d: PASS (result: %d)\n", i+1, result)
		} else {
			fmt.Printf("Test %d: FAIL (expected: %d, got: %d)\n", i+1, test.expected, result)
		}
	}
}
