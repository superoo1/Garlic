// Code generated from /Users/liwenda/Downloads/gengine-main/internal/iantlr/gengine.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 64, 565,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3,
	12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3,
	33, 3, 34, 3, 34, 5, 34, 263, 10, 34, 3, 34, 6, 34, 266, 10, 34, 13, 34,
	14, 34, 267, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3,
	36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3,
	45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47,
	3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3,
	49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51,
	3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53, 3, 53,
	3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 55, 3, 55, 3, 55, 3, 55, 3,
	55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 59, 6,
	59, 404, 10, 59, 13, 59, 14, 59, 405, 3, 59, 7, 59, 409, 10, 59, 12, 59,
	14, 59, 412, 11, 59, 3, 60, 6, 60, 415, 10, 60, 13, 60, 14, 60, 416, 3,
	61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65,
	3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69, 3, 69, 3,
	70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74,
	3, 74, 3, 74, 3, 75, 3, 75, 3, 75, 3, 76, 3, 76, 3, 76, 3, 77, 3, 77, 3,
	77, 3, 78, 3, 78, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3, 82,
	3, 83, 3, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 86, 3, 86, 3, 86, 3, 86, 3,
	86, 3, 86, 7, 86, 484, 10, 86, 12, 86, 14, 86, 487, 11, 86, 3, 86, 3, 86,
	3, 87, 3, 87, 3, 87, 3, 87, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88, 3,
	89, 6, 89, 502, 10, 89, 13, 89, 14, 89, 503, 5, 89, 506, 10, 89, 3, 89,
	3, 89, 6, 89, 510, 10, 89, 13, 89, 14, 89, 511, 3, 89, 6, 89, 515, 10,
	89, 13, 89, 14, 89, 516, 3, 89, 3, 89, 3, 89, 3, 89, 6, 89, 523, 10, 89,
	13, 89, 14, 89, 524, 5, 89, 527, 10, 89, 3, 89, 3, 89, 6, 89, 531, 10,
	89, 13, 89, 14, 89, 532, 3, 89, 3, 89, 3, 89, 6, 89, 538, 10, 89, 13, 89,
	14, 89, 539, 3, 89, 3, 89, 5, 89, 544, 10, 89, 3, 90, 3, 90, 3, 90, 3,
	90, 7, 90, 550, 10, 90, 12, 90, 14, 90, 553, 11, 90, 3, 90, 3, 90, 3, 90,
	3, 90, 3, 91, 6, 91, 560, 10, 91, 13, 91, 14, 91, 561, 3, 91, 3, 91, 3,
	551, 2, 92, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 2, 15, 2, 17, 2, 19, 2,
	21, 2, 23, 2, 25, 2, 27, 2, 29, 2, 31, 2, 33, 2, 35, 2, 37, 2, 39, 2, 41,
	2, 43, 2, 45, 2, 47, 2, 49, 2, 51, 2, 53, 2, 55, 2, 57, 2, 59, 2, 61, 2,
	63, 2, 65, 2, 67, 2, 69, 8, 71, 9, 73, 10, 75, 11, 77, 12, 79, 13, 81,
	14, 83, 15, 85, 16, 87, 17, 89, 18, 91, 19, 93, 20, 95, 21, 97, 22, 99,
	23, 101, 24, 103, 25, 105, 26, 107, 27, 109, 28, 111, 29, 113, 30, 115,
	31, 117, 32, 119, 33, 121, 34, 123, 35, 125, 36, 127, 37, 129, 38, 131,
	39, 133, 40, 135, 41, 137, 42, 139, 43, 141, 44, 143, 45, 145, 46, 147,
	47, 149, 48, 151, 49, 153, 50, 155, 51, 157, 52, 159, 53, 161, 54, 163,
	55, 165, 56, 167, 57, 169, 58, 171, 59, 173, 60, 175, 61, 177, 62, 179,
	63, 181, 64, 3, 2, 33, 3, 2, 50, 59, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68,
	100, 100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71,
	103, 103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74,
	106, 106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77,
	109, 109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80,
	112, 112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83,
	115, 115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86,
	118, 118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89,
	121, 121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92,
	124, 124, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97,
	99, 124, 4, 2, 36, 36, 94, 94, 5, 2, 11, 12, 15, 15, 34, 34, 2, 557, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2,
	2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2,
	2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2,
	2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3,
	2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2,
	105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2,
	2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119,
	3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2,
	2, 127, 3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3,
	2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2,
	141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3, 2,
	2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3, 2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155,
	3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 2, 159, 3, 2, 2, 2, 2, 161, 3, 2, 2, 2,
	2, 163, 3, 2, 2, 2, 2, 165, 3, 2, 2, 2, 2, 167, 3, 2, 2, 2, 2, 169, 3,
	2, 2, 2, 2, 171, 3, 2, 2, 2, 2, 173, 3, 2, 2, 2, 2, 175, 3, 2, 2, 2, 2,
	177, 3, 2, 2, 2, 2, 179, 3, 2, 2, 2, 2, 181, 3, 2, 2, 2, 3, 183, 3, 2,
	2, 2, 5, 185, 3, 2, 2, 2, 7, 191, 3, 2, 2, 2, 9, 195, 3, 2, 2, 2, 11, 201,
	3, 2, 2, 2, 13, 206, 3, 2, 2, 2, 15, 208, 3, 2, 2, 2, 17, 210, 3, 2, 2,
	2, 19, 212, 3, 2, 2, 2, 21, 214, 3, 2, 2, 2, 23, 216, 3, 2, 2, 2, 25, 218,
	3, 2, 2, 2, 27, 220, 3, 2, 2, 2, 29, 222, 3, 2, 2, 2, 31, 224, 3, 2, 2,
	2, 33, 226, 3, 2, 2, 2, 35, 228, 3, 2, 2, 2, 37, 230, 3, 2, 2, 2, 39, 232,
	3, 2, 2, 2, 41, 234, 3, 2, 2, 2, 43, 236, 3, 2, 2, 2, 45, 238, 3, 2, 2,
	2, 47, 240, 3, 2, 2, 2, 49, 242, 3, 2, 2, 2, 51, 244, 3, 2, 2, 2, 53, 246,
	3, 2, 2, 2, 55, 248, 3, 2, 2, 2, 57, 250, 3, 2, 2, 2, 59, 252, 3, 2, 2,
	2, 61, 254, 3, 2, 2, 2, 63, 256, 3, 2, 2, 2, 65, 258, 3, 2, 2, 2, 67, 260,
	3, 2, 2, 2, 69, 269, 3, 2, 2, 2, 71, 275, 3, 2, 2, 2, 73, 281, 3, 2, 2,
	2, 75, 285, 3, 2, 2, 2, 77, 292, 3, 2, 2, 2, 79, 299, 3, 2, 2, 2, 81, 304,
	3, 2, 2, 2, 83, 308, 3, 2, 2, 2, 85, 313, 3, 2, 2, 2, 87, 316, 3, 2, 2,
	2, 89, 319, 3, 2, 2, 2, 91, 324, 3, 2, 2, 2, 93, 327, 3, 2, 2, 2, 95, 332,
	3, 2, 2, 2, 97, 339, 3, 2, 2, 2, 99, 343, 3, 2, 2, 2, 101, 349, 3, 2, 2,
	2, 103, 358, 3, 2, 2, 2, 105, 367, 3, 2, 2, 2, 107, 372, 3, 2, 2, 2, 109,
	378, 3, 2, 2, 2, 111, 383, 3, 2, 2, 2, 113, 392, 3, 2, 2, 2, 115, 398,
	3, 2, 2, 2, 117, 403, 3, 2, 2, 2, 119, 414, 3, 2, 2, 2, 121, 418, 3, 2,
	2, 2, 123, 420, 3, 2, 2, 2, 125, 422, 3, 2, 2, 2, 127, 424, 3, 2, 2, 2,
	129, 426, 3, 2, 2, 2, 131, 429, 3, 2, 2, 2, 133, 431, 3, 2, 2, 2, 135,
	433, 3, 2, 2, 2, 137, 436, 3, 2, 2, 2, 139, 439, 3, 2, 2, 2, 141, 442,
	3, 2, 2, 2, 143, 444, 3, 2, 2, 2, 145, 447, 3, 2, 2, 2, 147, 449, 3, 2,
	2, 2, 149, 452, 3, 2, 2, 2, 151, 455, 3, 2, 2, 2, 153, 458, 3, 2, 2, 2,
	155, 461, 3, 2, 2, 2, 157, 463, 3, 2, 2, 2, 159, 465, 3, 2, 2, 2, 161,
	467, 3, 2, 2, 2, 163, 469, 3, 2, 2, 2, 165, 471, 3, 2, 2, 2, 167, 473,
	3, 2, 2, 2, 169, 475, 3, 2, 2, 2, 171, 477, 3, 2, 2, 2, 173, 490, 3, 2,
	2, 2, 175, 494, 3, 2, 2, 2, 177, 543, 3, 2, 2, 2, 179, 545, 3, 2, 2, 2,
	181, 559, 3, 2, 2, 2, 183, 184, 7, 46, 2, 2, 184, 4, 3, 2, 2, 2, 185, 186,
	7, 66, 2, 2, 186, 187, 7, 112, 2, 2, 187, 188, 7, 99, 2, 2, 188, 189, 7,
	111, 2, 2, 189, 190, 7, 103, 2, 2, 190, 6, 3, 2, 2, 2, 191, 192, 7, 66,
	2, 2, 192, 193, 7, 107, 2, 2, 193, 194, 7, 102, 2, 2, 194, 8, 3, 2, 2,
	2, 195, 196, 7, 66, 2, 2, 196, 197, 7, 102, 2, 2, 197, 198, 7, 103, 2,
	2, 198, 199, 7, 117, 2, 2, 199, 200, 7, 101, 2, 2, 200, 10, 3, 2, 2, 2,
	201, 202, 7, 66, 2, 2, 202, 203, 7, 117, 2, 2, 203, 204, 7, 99, 2, 2, 204,
	205, 7, 110, 2, 2, 205, 12, 3, 2, 2, 2, 206, 207, 9, 2, 2, 2, 207, 14,
	3, 2, 2, 2, 208, 209, 9, 3, 2, 2, 209, 16, 3, 2, 2, 2, 210, 211, 9, 4,
	2, 2, 211, 18, 3, 2, 2, 2, 212, 213, 9, 5, 2, 2, 213, 20, 3, 2, 2, 2, 214,
	215, 9, 6, 2, 2, 215, 22, 3, 2, 2, 2, 216, 217, 9, 7, 2, 2, 217, 24, 3,
	2, 2, 2, 218, 219, 9, 8, 2, 2, 219, 26, 3, 2, 2, 2, 220, 221, 9, 9, 2,
	2, 221, 28, 3, 2, 2, 2, 222, 223, 9, 10, 2, 2, 223, 30, 3, 2, 2, 2, 224,
	225, 9, 11, 2, 2, 225, 32, 3, 2, 2, 2, 226, 227, 9, 12, 2, 2, 227, 34,
	3, 2, 2, 2, 228, 229, 9, 13, 2, 2, 229, 36, 3, 2, 2, 2, 230, 231, 9, 14,
	2, 2, 231, 38, 3, 2, 2, 2, 232, 233, 9, 15, 2, 2, 233, 40, 3, 2, 2, 2,
	234, 235, 9, 16, 2, 2, 235, 42, 3, 2, 2, 2, 236, 237, 9, 17, 2, 2, 237,
	44, 3, 2, 2, 2, 238, 239, 9, 18, 2, 2, 239, 46, 3, 2, 2, 2, 240, 241, 9,
	19, 2, 2, 241, 48, 3, 2, 2, 2, 242, 243, 9, 20, 2, 2, 243, 50, 3, 2, 2,
	2, 244, 245, 9, 21, 2, 2, 245, 52, 3, 2, 2, 2, 246, 247, 9, 22, 2, 2, 247,
	54, 3, 2, 2, 2, 248, 249, 9, 23, 2, 2, 249, 56, 3, 2, 2, 2, 250, 251, 9,
	24, 2, 2, 251, 58, 3, 2, 2, 2, 252, 253, 9, 25, 2, 2, 253, 60, 3, 2, 2,
	2, 254, 255, 9, 26, 2, 2, 255, 62, 3, 2, 2, 2, 256, 257, 9, 27, 2, 2, 257,
	64, 3, 2, 2, 2, 258, 259, 9, 28, 2, 2, 259, 66, 3, 2, 2, 2, 260, 262, 9,
	7, 2, 2, 261, 263, 7, 47, 2, 2, 262, 261, 3, 2, 2, 2, 262, 263, 3, 2, 2,
	2, 263, 265, 3, 2, 2, 2, 264, 266, 5, 13, 7, 2, 265, 264, 3, 2, 2, 2, 266,
	267, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 68, 3,
	2, 2, 2, 269, 270, 5, 19, 10, 2, 270, 271, 5, 43, 22, 2, 271, 272, 5, 55,
	28, 2, 272, 273, 5, 41, 21, 2, 273, 274, 5, 53, 27, 2, 274, 70, 3, 2, 2,
	2, 275, 276, 5, 39, 20, 2, 276, 277, 5, 31, 16, 2, 277, 278, 5, 41, 21,
	2, 278, 279, 5, 53, 27, 2, 279, 280, 5, 23, 12, 2, 280, 72, 3, 2, 2, 2,
	281, 282, 5, 21, 11, 2, 282, 283, 5, 15, 8, 2, 283, 284, 5, 63, 32, 2,
	284, 74, 3, 2, 2, 2, 285, 286, 5, 51, 26, 2, 286, 287, 5, 23, 12, 2, 287,
	288, 5, 19, 10, 2, 288, 289, 5, 43, 22, 2, 289, 290, 5, 41, 21, 2, 290,
	291, 5, 21, 11, 2, 291, 76, 3, 2, 2, 2, 292, 293, 5, 59, 30, 2, 293, 294,
	5, 31, 16, 2, 294, 295, 5, 53, 27, 2, 295, 296, 5, 29, 15, 2, 296, 297,
	5, 31, 16, 2, 297, 298, 5, 41, 21, 2, 298, 78, 3, 2, 2, 2, 299, 300, 5,
	49, 25, 2, 300, 301, 5, 23, 12, 2, 301, 302, 5, 27, 14, 2, 302, 303, 5,
	61, 31, 2, 303, 80, 3, 2, 2, 2, 304, 305, 5, 41, 21, 2, 305, 306, 5, 31,
	16, 2, 306, 307, 5, 37, 19, 2, 307, 82, 3, 2, 2, 2, 308, 309, 5, 49, 25,
	2, 309, 310, 5, 55, 28, 2, 310, 311, 5, 37, 19, 2, 311, 312, 5, 23, 12,
	2, 312, 84, 3, 2, 2, 2, 313, 314, 7, 40, 2, 2, 314, 315, 7, 40, 2, 2, 315,
	86, 3, 2, 2, 2, 316, 317, 7, 126, 2, 2, 317, 318, 7, 126, 2, 2, 318, 88,
	3, 2, 2, 2, 319, 320, 5, 19, 10, 2, 320, 321, 5, 43, 22, 2, 321, 322, 5,
	41, 21, 2, 322, 323, 5, 19, 10, 2, 323, 90, 3, 2, 2, 2, 324, 325, 5, 31,
	16, 2, 325, 326, 5, 25, 13, 2, 326, 92, 3, 2, 2, 2, 327, 328, 5, 23, 12,
	2, 328, 329, 5, 37, 19, 2, 329, 330, 5, 51, 26, 2, 330, 331, 5, 23, 12,
	2, 331, 94, 3, 2, 2, 2, 332, 333, 5, 49, 25, 2, 333, 334, 5, 23, 12, 2,
	334, 335, 5, 53, 27, 2, 335, 336, 5, 55, 28, 2, 336, 337, 5, 49, 25, 2,
	337, 338, 5, 41, 21, 2, 338, 96, 3, 2, 2, 2, 339, 340, 5, 25, 13, 2, 340,
	341, 5, 43, 22, 2, 341, 342, 5, 49, 25, 2, 342, 98, 3, 2, 2, 2, 343, 344,
	5, 17, 9, 2, 344, 345, 5, 49, 25, 2, 345, 346, 5, 23, 12, 2, 346, 347,
	5, 15, 8, 2, 347, 348, 5, 35, 18, 2, 348, 100, 3, 2, 2, 2, 349, 350, 5,
	25, 13, 2, 350, 351, 5, 43, 22, 2, 351, 352, 5, 49, 25, 2, 352, 353, 5,
	49, 25, 2, 353, 354, 5, 15, 8, 2, 354, 355, 5, 41, 21, 2, 355, 356, 5,
	27, 14, 2, 356, 357, 5, 23, 12, 2, 357, 102, 3, 2, 2, 2, 358, 359, 5, 19,
	10, 2, 359, 360, 5, 43, 22, 2, 360, 361, 5, 41, 21, 2, 361, 362, 5, 53,
	27, 2, 362, 363, 5, 31, 16, 2, 363, 364, 5, 41, 21, 2, 364, 365, 5, 55,
	28, 2, 365, 366, 5, 23, 12, 2, 366, 104, 3, 2, 2, 2, 367, 368, 5, 53, 27,
	2, 368, 369, 5, 49, 25, 2, 369, 370, 5, 55, 28, 2, 370, 371, 5, 23, 12,
	2, 371, 106, 3, 2, 2, 2, 372, 373, 5, 25, 13, 2, 373, 374, 5, 15, 8, 2,
	374, 375, 5, 37, 19, 2, 375, 376, 5, 51, 26, 2, 376, 377, 5, 23, 12, 2,
	377, 108, 3, 2, 2, 2, 378, 379, 5, 41, 21, 2, 379, 380, 5, 55, 28, 2, 380,
	381, 5, 37, 19, 2, 381, 382, 5, 37, 19, 2, 382, 110, 3, 2, 2, 2, 383, 384,
	5, 51, 26, 2, 384, 385, 5, 15, 8, 2, 385, 386, 5, 37, 19, 2, 386, 387,
	5, 31, 16, 2, 387, 388, 5, 23, 12, 2, 388, 389, 5, 41, 21, 2, 389, 390,
	5, 19, 10, 2, 390, 391, 5, 23, 12, 2, 391, 112, 3, 2, 2, 2, 392, 393, 5,
	17, 9, 2, 393, 394, 5, 23, 12, 2, 394, 395, 5, 27, 14, 2, 395, 396, 5,
	31, 16, 2, 396, 397, 5, 41, 21, 2, 397, 114, 3, 2, 2, 2, 398, 399, 5, 23,
	12, 2, 399, 400, 5, 41, 21, 2, 400, 401, 5, 21, 11, 2, 401, 116, 3, 2,
	2, 2, 402, 404, 9, 29, 2, 2, 403, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2,
	405, 403, 3, 2, 2, 2, 405, 406, 3, 2, 2, 2, 406, 410, 3, 2, 2, 2, 407,
	409, 9, 30, 2, 2, 408, 407, 3, 2, 2, 2, 409, 412, 3, 2, 2, 2, 410, 408,
	3, 2, 2, 2, 410, 411, 3, 2, 2, 2, 411, 118, 3, 2, 2, 2, 412, 410, 3, 2,
	2, 2, 413, 415, 4, 50, 59, 2, 414, 413, 3, 2, 2, 2, 415, 416, 3, 2, 2,
	2, 416, 414, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 120, 3, 2, 2, 2, 418,
	419, 7, 45, 2, 2, 419, 122, 3, 2, 2, 2, 420, 421, 7, 47, 2, 2, 421, 124,
	3, 2, 2, 2, 422, 423, 7, 49, 2, 2, 423, 126, 3, 2, 2, 2, 424, 425, 7, 44,
	2, 2, 425, 128, 3, 2, 2, 2, 426, 427, 7, 63, 2, 2, 427, 428, 7, 63, 2,
	2, 428, 130, 3, 2, 2, 2, 429, 430, 7, 64, 2, 2, 430, 132, 3, 2, 2, 2, 431,
	432, 7, 62, 2, 2, 432, 134, 3, 2, 2, 2, 433, 434, 7, 64, 2, 2, 434, 435,
	7, 63, 2, 2, 435, 136, 3, 2, 2, 2, 436, 437, 7, 62, 2, 2, 437, 438, 7,
	63, 2, 2, 438, 138, 3, 2, 2, 2, 439, 440, 7, 35, 2, 2, 440, 441, 7, 63,
	2, 2, 441, 140, 3, 2, 2, 2, 442, 443, 7, 35, 2, 2, 443, 142, 3, 2, 2, 2,
	444, 445, 7, 60, 2, 2, 445, 446, 7, 63, 2, 2, 446, 144, 3, 2, 2, 2, 447,
	448, 7, 63, 2, 2, 448, 146, 3, 2, 2, 2, 449, 450, 7, 45, 2, 2, 450, 451,
	7, 63, 2, 2, 451, 148, 3, 2, 2, 2, 452, 453, 7, 47, 2, 2, 453, 454, 7,
	63, 2, 2, 454, 150, 3, 2, 2, 2, 455, 456, 7, 44, 2, 2, 456, 457, 7, 63,
	2, 2, 457, 152, 3, 2, 2, 2, 458, 459, 7, 49, 2, 2, 459, 460, 7, 63, 2,
	2, 460, 154, 3, 2, 2, 2, 461, 462, 7, 93, 2, 2, 462, 156, 3, 2, 2, 2, 463,
	464, 7, 95, 2, 2, 464, 158, 3, 2, 2, 2, 465, 466, 7, 61, 2, 2, 466, 160,
	3, 2, 2, 2, 467, 468, 7, 125, 2, 2, 468, 162, 3, 2, 2, 2, 469, 470, 7,
	127, 2, 2, 470, 164, 3, 2, 2, 2, 471, 472, 7, 42, 2, 2, 472, 166, 3, 2,
	2, 2, 473, 474, 7, 43, 2, 2, 474, 168, 3, 2, 2, 2, 475, 476, 7, 48, 2,
	2, 476, 170, 3, 2, 2, 2, 477, 485, 7, 36, 2, 2, 478, 479, 7, 94, 2, 2,
	479, 484, 11, 2, 2, 2, 480, 481, 7, 36, 2, 2, 481, 484, 7, 36, 2, 2, 482,
	484, 10, 31, 2, 2, 483, 478, 3, 2, 2, 2, 483, 480, 3, 2, 2, 2, 483, 482,
	3, 2, 2, 2, 484, 487, 3, 2, 2, 2, 485, 483, 3, 2, 2, 2, 485, 486, 3, 2,
	2, 2, 486, 488, 3, 2, 2, 2, 487, 485, 3, 2, 2, 2, 488, 489, 7, 36, 2, 2,
	489, 172, 3, 2, 2, 2, 490, 491, 5, 117, 59, 2, 491, 492, 5, 169, 85, 2,
	492, 493, 5, 117, 59, 2, 493, 174, 3, 2, 2, 2, 494, 495, 5, 117, 59, 2,
	495, 496, 5, 169, 85, 2, 496, 497, 5, 117, 59, 2, 497, 498, 5, 169, 85,
	2, 498, 499, 5, 117, 59, 2, 499, 176, 3, 2, 2, 2, 500, 502, 5, 13, 7, 2,
	501, 500, 3, 2, 2, 2, 502, 503, 3, 2, 2, 2, 503, 501, 3, 2, 2, 2, 503,
	504, 3, 2, 2, 2, 504, 506, 3, 2, 2, 2, 505, 501, 3, 2, 2, 2, 505, 506,
	3, 2, 2, 2, 506, 507, 3, 2, 2, 2, 507, 509, 7, 48, 2, 2, 508, 510, 5, 13,
	7, 2, 509, 508, 3, 2, 2, 2, 510, 511, 3, 2, 2, 2, 511, 509, 3, 2, 2, 2,
	511, 512, 3, 2, 2, 2, 512, 544, 3, 2, 2, 2, 513, 515, 5, 13, 7, 2, 514,
	513, 3, 2, 2, 2, 515, 516, 3, 2, 2, 2, 516, 514, 3, 2, 2, 2, 516, 517,
	3, 2, 2, 2, 517, 518, 3, 2, 2, 2, 518, 519, 7, 48, 2, 2, 519, 520, 5, 67,
	34, 2, 520, 544, 3, 2, 2, 2, 521, 523, 5, 13, 7, 2, 522, 521, 3, 2, 2,
	2, 523, 524, 3, 2, 2, 2, 524, 522, 3, 2, 2, 2, 524, 525, 3, 2, 2, 2, 525,
	527, 3, 2, 2, 2, 526, 522, 3, 2, 2, 2, 526, 527, 3, 2, 2, 2, 527, 528,
	3, 2, 2, 2, 528, 530, 7, 48, 2, 2, 529, 531, 5, 13, 7, 2, 530, 529, 3,
	2, 2, 2, 531, 532, 3, 2, 2, 2, 532, 530, 3, 2, 2, 2, 532, 533, 3, 2, 2,
	2, 533, 534, 3, 2, 2, 2, 534, 535, 5, 67, 34, 2, 535, 544, 3, 2, 2, 2,
	536, 538, 5, 13, 7, 2, 537, 536, 3, 2, 2, 2, 538, 539, 3, 2, 2, 2, 539,
	537, 3, 2, 2, 2, 539, 540, 3, 2, 2, 2, 540, 541, 3, 2, 2, 2, 541, 542,
	5, 67, 34, 2, 542, 544, 3, 2, 2, 2, 543, 505, 3, 2, 2, 2, 543, 514, 3,
	2, 2, 2, 543, 526, 3, 2, 2, 2, 543, 537, 3, 2, 2, 2, 544, 178, 3, 2, 2,
	2, 545, 546, 7, 49, 2, 2, 546, 547, 7, 49, 2, 2, 547, 551, 3, 2, 2, 2,
	548, 550, 11, 2, 2, 2, 549, 548, 3, 2, 2, 2, 550, 553, 3, 2, 2, 2, 551,
	552, 3, 2, 2, 2, 551, 549, 3, 2, 2, 2, 552, 554, 3, 2, 2, 2, 553, 551,
	3, 2, 2, 2, 554, 555, 7, 12, 2, 2, 555, 556, 3, 2, 2, 2, 556, 557, 8, 90,
	2, 2, 557, 180, 3, 2, 2, 2, 558, 560, 9, 32, 2, 2, 559, 558, 3, 2, 2, 2,
	560, 561, 3, 2, 2, 2, 561, 559, 3, 2, 2, 2, 561, 562, 3, 2, 2, 2, 562,
	563, 3, 2, 2, 2, 563, 564, 8, 91, 2, 2, 564, 182, 3, 2, 2, 2, 22, 2, 262,
	267, 405, 408, 410, 416, 483, 485, 503, 505, 511, 516, 524, 526, 532, 539,
	543, 551, 561, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'@name'", "'@id'", "'@desc'", "'@sal'", "", "", "", "", "",
	"", "", "", "'&&'", "'||'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "'+'", "'-'", "'/'", "'*'", "'=='", "'>'", "'<'", "'>='",
	"'<='", "'!='", "'!'", "':='", "'='", "'+='", "'-='", "'*='", "'/='", "'['",
	"']'", "';'", "'{'", "'}'", "'('", "')'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "COUNT", "MINTE", "DAY", "SECOND", "WITHIN", "REGX",
	"NIL", "RULE", "AND", "OR", "CONC", "IF", "ELSE", "RETURN", "FOR", "BREAK",
	"FORRANGE", "CONTINUE", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN",
	"END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT",
	"LT", "GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL",
	"MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE",
	"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME",
	"DOUBLEDOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "DEC_DIGIT", "A", "B", "C", "D",
	"E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S",
	"T", "U", "V", "W", "X", "Y", "Z", "EXPONENT_NUM_PART", "COUNT", "MINTE",
	"DAY", "SECOND", "WITHIN", "REGX", "NIL", "RULE", "AND", "OR", "CONC",
	"IF", "ELSE", "RETURN", "FOR", "BREAK", "FORRANGE", "CONTINUE", "TRUE",
	"FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN", "END", "SIMPLENAME", "INT",
	"PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT", "LT", "GTE", "LTE", "NOTEQUALS",
	"NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL", "MULTIEQUAL", "DIVEQUAL",
	"LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE", "RR_BRACE", "LR_BRACKET",
	"RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME", "DOUBLEDOTTEDNAME",
	"REAL_LITERAL", "SL_COMMENT", "WS",
}

type gengineLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewgengineLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *gengineLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgengineLexer(input antlr.CharStream) *gengineLexer {
	l := new(gengineLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "gengine.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// gengineLexer tokens.
const (
	gengineLexerT__0             = 1
	gengineLexerT__1             = 2
	gengineLexerT__2             = 3
	gengineLexerT__3             = 4
	gengineLexerT__4             = 5
	gengineLexerCOUNT            = 6
	gengineLexerMINTE            = 7
	gengineLexerDAY              = 8
	gengineLexerSECOND           = 9
	gengineLexerWITHIN           = 10
	gengineLexerREGX             = 11
	gengineLexerNIL              = 12
	gengineLexerRULE             = 13
	gengineLexerAND              = 14
	gengineLexerOR               = 15
	gengineLexerCONC             = 16
	gengineLexerIF               = 17
	gengineLexerELSE             = 18
	gengineLexerRETURN           = 19
	gengineLexerFOR              = 20
	gengineLexerBREAK            = 21
	gengineLexerFORRANGE         = 22
	gengineLexerCONTINUE         = 23
	gengineLexerTRUE             = 24
	gengineLexerFALSE            = 25
	gengineLexerNULL_LITERAL     = 26
	gengineLexerSALIENCE         = 27
	gengineLexerBEGIN            = 28
	gengineLexerEND              = 29
	gengineLexerSIMPLENAME       = 30
	gengineLexerINT              = 31
	gengineLexerPLUS             = 32
	gengineLexerMINUS            = 33
	gengineLexerDIV              = 34
	gengineLexerMUL              = 35
	gengineLexerEQUALS           = 36
	gengineLexerGT               = 37
	gengineLexerLT               = 38
	gengineLexerGTE              = 39
	gengineLexerLTE              = 40
	gengineLexerNOTEQUALS        = 41
	gengineLexerNOT              = 42
	gengineLexerASSIGN           = 43
	gengineLexerSET              = 44
	gengineLexerPLUSEQUAL        = 45
	gengineLexerMINUSEQUAL       = 46
	gengineLexerMULTIEQUAL       = 47
	gengineLexerDIVEQUAL         = 48
	gengineLexerLSQARE           = 49
	gengineLexerRSQARE           = 50
	gengineLexerSEMICOLON        = 51
	gengineLexerLR_BRACE         = 52
	gengineLexerRR_BRACE         = 53
	gengineLexerLR_BRACKET       = 54
	gengineLexerRR_BRACKET       = 55
	gengineLexerDOT              = 56
	gengineLexerDQUOTA_STRING    = 57
	gengineLexerDOTTEDNAME       = 58
	gengineLexerDOUBLEDOTTEDNAME = 59
	gengineLexerREAL_LITERAL     = 60
	gengineLexerSL_COMMENT       = 61
	gengineLexerWS               = 62
)
