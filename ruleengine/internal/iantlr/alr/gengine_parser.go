// Code generated from /Users/liwenda/Downloads/gengine-main/internal/iantlr/gengine.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // gengine

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 64, 398,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 3,
	2, 6, 2, 100, 10, 2, 13, 2, 14, 2, 101, 3, 3, 3, 3, 3, 3, 5, 3, 107, 10,
	3, 3, 3, 5, 3, 110, 10, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 7, 8, 126, 10, 8, 12, 8, 14, 8,
	129, 11, 8, 3, 8, 5, 8, 132, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 145, 10, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 7, 10, 153, 10, 10, 12, 10, 14, 10, 156, 11, 10, 3, 10,
	3, 10, 3, 11, 3, 11, 3, 11, 5, 11, 163, 10, 11, 3, 11, 3, 11, 5, 11, 167,
	10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 173, 10, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 183, 10, 11, 12, 11, 14,
	11, 186, 11, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 194,
	10, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12,
	204, 10, 12, 12, 12, 14, 12, 207, 11, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 5, 13, 215, 10, 13, 3, 14, 3, 14, 5, 14, 219, 10, 14, 3, 14,
	3, 14, 3, 14, 5, 14, 224, 10, 14, 3, 15, 3, 15, 5, 15, 228, 10, 15, 3,
	16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 246, 10, 17, 12, 17, 14, 17,
	249, 11, 17, 3, 17, 5, 17, 252, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 296, 10, 24, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 305, 10, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 5, 25, 315, 10, 25, 7, 25, 317,
	10, 25, 12, 25, 14, 25, 320, 11, 25, 3, 26, 5, 26, 323, 10, 26, 3, 26,
	3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 5, 30, 334, 10,
	30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 5, 33,
	345, 10, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 5, 34, 352, 10, 34, 3,
	34, 3, 34, 3, 35, 3, 35, 3, 35, 5, 35, 359, 10, 35, 3, 35, 3, 35, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 38, 3, 38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3,
	41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 45, 5, 45, 386, 10, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3,
	48, 3, 48, 3, 49, 3, 49, 3, 49, 2, 4, 20, 22, 50, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86,
	88, 90, 92, 94, 96, 2, 10, 3, 2, 9, 11, 3, 2, 26, 27, 4, 2, 32, 32, 60,
	61, 3, 2, 34, 35, 3, 2, 36, 37, 3, 2, 38, 43, 3, 2, 16, 17, 3, 2, 45, 50,
	2, 414, 2, 99, 3, 2, 2, 2, 4, 103, 3, 2, 2, 2, 6, 115, 3, 2, 2, 2, 8, 117,
	3, 2, 2, 2, 10, 119, 3, 2, 2, 2, 12, 122, 3, 2, 2, 2, 14, 127, 3, 2, 2,
	2, 16, 144, 3, 2, 2, 2, 18, 146, 3, 2, 2, 2, 20, 172, 3, 2, 2, 2, 22, 193,
	3, 2, 2, 2, 24, 214, 3, 2, 2, 2, 26, 218, 3, 2, 2, 2, 28, 225, 3, 2, 2,
	2, 30, 229, 3, 2, 2, 2, 32, 239, 3, 2, 2, 2, 34, 253, 3, 2, 2, 2, 36, 260,
	3, 2, 2, 2, 38, 265, 3, 2, 2, 2, 40, 275, 3, 2, 2, 2, 42, 277, 3, 2, 2,
	2, 44, 285, 3, 2, 2, 2, 46, 295, 3, 2, 2, 2, 48, 304, 3, 2, 2, 2, 50, 322,
	3, 2, 2, 2, 52, 326, 3, 2, 2, 2, 54, 328, 3, 2, 2, 2, 56, 330, 3, 2, 2,
	2, 58, 333, 3, 2, 2, 2, 60, 337, 3, 2, 2, 2, 62, 339, 3, 2, 2, 2, 64, 341,
	3, 2, 2, 2, 66, 348, 3, 2, 2, 2, 68, 355, 3, 2, 2, 2, 70, 362, 3, 2, 2,
	2, 72, 364, 3, 2, 2, 2, 74, 366, 3, 2, 2, 2, 76, 368, 3, 2, 2, 2, 78, 370,
	3, 2, 2, 2, 80, 372, 3, 2, 2, 2, 82, 374, 3, 2, 2, 2, 84, 376, 3, 2, 2,
	2, 86, 378, 3, 2, 2, 2, 88, 380, 3, 2, 2, 2, 90, 389, 3, 2, 2, 2, 92, 391,
	3, 2, 2, 2, 94, 393, 3, 2, 2, 2, 96, 395, 3, 2, 2, 2, 98, 100, 5, 4, 3,
	2, 99, 98, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 99, 3, 2, 2, 2, 101,
	102, 3, 2, 2, 2, 102, 3, 3, 2, 2, 2, 103, 104, 7, 15, 2, 2, 104, 106, 5,
	6, 4, 2, 105, 107, 5, 8, 5, 2, 106, 105, 3, 2, 2, 2, 106, 107, 3, 2, 2,
	2, 107, 109, 3, 2, 2, 2, 108, 110, 5, 10, 6, 2, 109, 108, 3, 2, 2, 2, 109,
	110, 3, 2, 2, 2, 110, 111, 3, 2, 2, 2, 111, 112, 7, 30, 2, 2, 112, 113,
	5, 12, 7, 2, 113, 114, 7, 31, 2, 2, 114, 5, 3, 2, 2, 2, 115, 116, 5, 60,
	31, 2, 116, 7, 3, 2, 2, 2, 117, 118, 5, 60, 31, 2, 118, 9, 3, 2, 2, 2,
	119, 120, 7, 29, 2, 2, 120, 121, 5, 50, 26, 2, 121, 11, 3, 2, 2, 2, 122,
	123, 5, 14, 8, 2, 123, 13, 3, 2, 2, 2, 124, 126, 5, 16, 9, 2, 125, 124,
	3, 2, 2, 2, 126, 129, 3, 2, 2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2,
	2, 2, 128, 131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 130, 132, 5, 28, 15,
	2, 131, 130, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 15, 3, 2, 2, 2, 133,
	145, 5, 32, 17, 2, 134, 145, 5, 64, 33, 2, 135, 145, 5, 66, 34, 2, 136,
	145, 5, 68, 35, 2, 137, 145, 5, 26, 14, 2, 138, 145, 5, 18, 10, 2, 139,
	145, 5, 38, 20, 2, 140, 145, 5, 40, 21, 2, 141, 145, 5, 42, 22, 2, 142,
	145, 5, 44, 23, 2, 143, 145, 5, 30, 16, 2, 144, 133, 3, 2, 2, 2, 144, 134,
	3, 2, 2, 2, 144, 135, 3, 2, 2, 2, 144, 136, 3, 2, 2, 2, 144, 137, 3, 2,
	2, 2, 144, 138, 3, 2, 2, 2, 144, 139, 3, 2, 2, 2, 144, 140, 3, 2, 2, 2,
	144, 141, 3, 2, 2, 2, 144, 142, 3, 2, 2, 2, 144, 143, 3, 2, 2, 2, 145,
	17, 3, 2, 2, 2, 146, 147, 7, 18, 2, 2, 147, 154, 7, 54, 2, 2, 148, 153,
	5, 64, 33, 2, 149, 153, 5, 66, 34, 2, 150, 153, 5, 68, 35, 2, 151, 153,
	5, 26, 14, 2, 152, 148, 3, 2, 2, 2, 152, 149, 3, 2, 2, 2, 152, 150, 3,
	2, 2, 2, 152, 151, 3, 2, 2, 2, 153, 156, 3, 2, 2, 2, 154, 152, 3, 2, 2,
	2, 154, 155, 3, 2, 2, 2, 155, 157, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2, 157,
	158, 7, 55, 2, 2, 158, 19, 3, 2, 2, 2, 159, 160, 8, 11, 1, 2, 160, 173,
	5, 22, 12, 2, 161, 163, 5, 86, 44, 2, 162, 161, 3, 2, 2, 2, 162, 163, 3,
	2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 173, 5, 24, 13, 2, 165, 167, 5, 86,
	44, 2, 166, 165, 3, 2, 2, 2, 166, 167, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2,
	168, 169, 7, 56, 2, 2, 169, 170, 5, 20, 11, 2, 170, 171, 7, 57, 2, 2, 171,
	173, 3, 2, 2, 2, 172, 159, 3, 2, 2, 2, 172, 162, 3, 2, 2, 2, 172, 166,
	3, 2, 2, 2, 173, 184, 3, 2, 2, 2, 174, 175, 12, 6, 2, 2, 175, 176, 5, 76,
	39, 2, 176, 177, 5, 20, 11, 7, 177, 183, 3, 2, 2, 2, 178, 179, 12, 5, 2,
	2, 179, 180, 5, 80, 41, 2, 180, 181, 5, 20, 11, 6, 181, 183, 3, 2, 2, 2,
	182, 174, 3, 2, 2, 2, 182, 178, 3, 2, 2, 2, 183, 186, 3, 2, 2, 2, 184,
	182, 3, 2, 2, 2, 184, 185, 3, 2, 2, 2, 185, 21, 3, 2, 2, 2, 186, 184, 3,
	2, 2, 2, 187, 188, 8, 12, 1, 2, 188, 194, 5, 24, 13, 2, 189, 190, 7, 56,
	2, 2, 190, 191, 5, 22, 12, 2, 191, 192, 7, 57, 2, 2, 192, 194, 3, 2, 2,
	2, 193, 187, 3, 2, 2, 2, 193, 189, 3, 2, 2, 2, 194, 205, 3, 2, 2, 2, 195,
	196, 12, 6, 2, 2, 196, 197, 5, 74, 38, 2, 197, 198, 5, 22, 12, 7, 198,
	204, 3, 2, 2, 2, 199, 200, 12, 5, 2, 2, 200, 201, 5, 72, 37, 2, 201, 202,
	5, 22, 12, 6, 202, 204, 3, 2, 2, 2, 203, 195, 3, 2, 2, 2, 203, 199, 3,
	2, 2, 2, 204, 207, 3, 2, 2, 2, 205, 203, 3, 2, 2, 2, 205, 206, 3, 2, 2,
	2, 206, 23, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 208, 215, 5, 64, 33, 2, 209,
	215, 5, 66, 34, 2, 210, 215, 5, 68, 35, 2, 211, 215, 5, 46, 24, 2, 212,
	215, 5, 88, 45, 2, 213, 215, 5, 70, 36, 2, 214, 208, 3, 2, 2, 2, 214, 209,
	3, 2, 2, 2, 214, 210, 3, 2, 2, 2, 214, 211, 3, 2, 2, 2, 214, 212, 3, 2,
	2, 2, 214, 213, 3, 2, 2, 2, 215, 25, 3, 2, 2, 2, 216, 219, 5, 88, 45, 2,
	217, 219, 5, 70, 36, 2, 218, 216, 3, 2, 2, 2, 218, 217, 3, 2, 2, 2, 219,
	220, 3, 2, 2, 2, 220, 223, 5, 82, 42, 2, 221, 224, 5, 22, 12, 2, 222, 224,
	5, 20, 11, 2, 223, 221, 3, 2, 2, 2, 223, 222, 3, 2, 2, 2, 224, 27, 3, 2,
	2, 2, 225, 227, 7, 21, 2, 2, 226, 228, 5, 20, 11, 2, 227, 226, 3, 2, 2,
	2, 227, 228, 3, 2, 2, 2, 228, 29, 3, 2, 2, 2, 229, 230, 7, 8, 2, 2, 230,
	231, 7, 56, 2, 2, 231, 232, 5, 20, 11, 2, 232, 233, 7, 57, 2, 2, 233, 234,
	5, 78, 40, 2, 234, 235, 5, 56, 29, 2, 235, 236, 7, 12, 2, 2, 236, 237,
	5, 54, 28, 2, 237, 238, 5, 52, 27, 2, 238, 31, 3, 2, 2, 2, 239, 240, 7,
	19, 2, 2, 240, 241, 5, 20, 11, 2, 241, 242, 7, 54, 2, 2, 242, 243, 5, 14,
	8, 2, 243, 247, 7, 55, 2, 2, 244, 246, 5, 34, 18, 2, 245, 244, 3, 2, 2,
	2, 246, 249, 3, 2, 2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248,
	251, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 250, 252, 5, 36, 19, 2, 251, 250,
	3, 2, 2, 2, 251, 252, 3, 2, 2, 2, 252, 33, 3, 2, 2, 2, 253, 254, 7, 20,
	2, 2, 254, 255, 7, 19, 2, 2, 255, 256, 5, 20, 11, 2, 256, 257, 7, 54, 2,
	2, 257, 258, 5, 14, 8, 2, 258, 259, 7, 55, 2, 2, 259, 35, 3, 2, 2, 2, 260,
	261, 7, 20, 2, 2, 261, 262, 7, 54, 2, 2, 262, 263, 5, 14, 8, 2, 263, 264,
	7, 55, 2, 2, 264, 37, 3, 2, 2, 2, 265, 266, 7, 22, 2, 2, 266, 267, 5, 26,
	14, 2, 267, 268, 7, 53, 2, 2, 268, 269, 5, 20, 11, 2, 269, 270, 7, 53,
	2, 2, 270, 271, 5, 26, 14, 2, 271, 272, 7, 54, 2, 2, 272, 273, 5, 14, 8,
	2, 273, 274, 7, 55, 2, 2, 274, 39, 3, 2, 2, 2, 275, 276, 7, 23, 2, 2, 276,
	41, 3, 2, 2, 2, 277, 278, 7, 24, 2, 2, 278, 279, 5, 70, 36, 2, 279, 280,
	5, 84, 43, 2, 280, 281, 5, 70, 36, 2, 281, 282, 7, 54, 2, 2, 282, 283,
	5, 14, 8, 2, 283, 284, 7, 55, 2, 2, 284, 43, 3, 2, 2, 2, 285, 286, 7, 25,
	2, 2, 286, 45, 3, 2, 2, 2, 287, 296, 5, 62, 32, 2, 288, 296, 5, 50, 26,
	2, 289, 296, 5, 58, 30, 2, 290, 296, 5, 60, 31, 2, 291, 296, 5, 90, 46,
	2, 292, 296, 5, 92, 47, 2, 293, 296, 5, 94, 48, 2, 294, 296, 5, 96, 49,
	2, 295, 287, 3, 2, 2, 2, 295, 288, 3, 2, 2, 2, 295, 289, 3, 2, 2, 2, 295,
	290, 3, 2, 2, 2, 295, 291, 3, 2, 2, 2, 295, 292, 3, 2, 2, 2, 295, 293,
	3, 2, 2, 2, 295, 294, 3, 2, 2, 2, 296, 47, 3, 2, 2, 2, 297, 305, 5, 46,
	24, 2, 298, 305, 5, 70, 36, 2, 299, 305, 5, 64, 33, 2, 300, 305, 5, 66,
	34, 2, 301, 305, 5, 68, 35, 2, 302, 305, 5, 88, 45, 2, 303, 305, 5, 20,
	11, 2, 304, 297, 3, 2, 2, 2, 304, 298, 3, 2, 2, 2, 304, 299, 3, 2, 2, 2,
	304, 300, 3, 2, 2, 2, 304, 301, 3, 2, 2, 2, 304, 302, 3, 2, 2, 2, 304,
	303, 3, 2, 2, 2, 305, 318, 3, 2, 2, 2, 306, 314, 7, 3, 2, 2, 307, 315,
	5, 46, 24, 2, 308, 315, 5, 70, 36, 2, 309, 315, 5, 64, 33, 2, 310, 315,
	5, 66, 34, 2, 311, 315, 5, 68, 35, 2, 312, 315, 5, 88, 45, 2, 313, 315,
	5, 20, 11, 2, 314, 307, 3, 2, 2, 2, 314, 308, 3, 2, 2, 2, 314, 309, 3,
	2, 2, 2, 314, 310, 3, 2, 2, 2, 314, 311, 3, 2, 2, 2, 314, 312, 3, 2, 2,
	2, 314, 313, 3, 2, 2, 2, 315, 317, 3, 2, 2, 2, 316, 306, 3, 2, 2, 2, 317,
	320, 3, 2, 2, 2, 318, 316, 3, 2, 2, 2, 318, 319, 3, 2, 2, 2, 319, 49, 3,
	2, 2, 2, 320, 318, 3, 2, 2, 2, 321, 323, 7, 35, 2, 2, 322, 321, 3, 2, 2,
	2, 322, 323, 3, 2, 2, 2, 323, 324, 3, 2, 2, 2, 324, 325, 7, 33, 2, 2, 325,
	51, 3, 2, 2, 2, 326, 327, 9, 2, 2, 2, 327, 53, 3, 2, 2, 2, 328, 329, 7,
	33, 2, 2, 329, 55, 3, 2, 2, 2, 330, 331, 7, 33, 2, 2, 331, 57, 3, 2, 2,
	2, 332, 334, 7, 35, 2, 2, 333, 332, 3, 2, 2, 2, 333, 334, 3, 2, 2, 2, 334,
	335, 3, 2, 2, 2, 335, 336, 7, 62, 2, 2, 336, 59, 3, 2, 2, 2, 337, 338,
	7, 59, 2, 2, 338, 61, 3, 2, 2, 2, 339, 340, 9, 3, 2, 2, 340, 63, 3, 2,
	2, 2, 341, 342, 7, 32, 2, 2, 342, 344, 7, 56, 2, 2, 343, 345, 5, 48, 25,
	2, 344, 343, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346,
	347, 7, 57, 2, 2, 347, 65, 3, 2, 2, 2, 348, 349, 7, 60, 2, 2, 349, 351,
	7, 56, 2, 2, 350, 352, 5, 48, 25, 2, 351, 350, 3, 2, 2, 2, 351, 352, 3,
	2, 2, 2, 352, 353, 3, 2, 2, 2, 353, 354, 7, 57, 2, 2, 354, 67, 3, 2, 2,
	2, 355, 356, 7, 61, 2, 2, 356, 358, 7, 56, 2, 2, 357, 359, 5, 48, 25, 2,
	358, 357, 3, 2, 2, 2, 358, 359, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360,
	361, 7, 57, 2, 2, 361, 69, 3, 2, 2, 2, 362, 363, 9, 4, 2, 2, 363, 71, 3,
	2, 2, 2, 364, 365, 9, 5, 2, 2, 365, 73, 3, 2, 2, 2, 366, 367, 9, 6, 2,
	2, 367, 75, 3, 2, 2, 2, 368, 369, 9, 7, 2, 2, 369, 77, 3, 2, 2, 2, 370,
	371, 9, 7, 2, 2, 371, 79, 3, 2, 2, 2, 372, 373, 9, 8, 2, 2, 373, 81, 3,
	2, 2, 2, 374, 375, 9, 9, 2, 2, 375, 83, 3, 2, 2, 2, 376, 377, 7, 45, 2,
	2, 377, 85, 3, 2, 2, 2, 378, 379, 7, 44, 2, 2, 379, 87, 3, 2, 2, 2, 380,
	381, 5, 70, 36, 2, 381, 385, 7, 51, 2, 2, 382, 386, 5, 50, 26, 2, 383,
	386, 5, 60, 31, 2, 384, 386, 5, 70, 36, 2, 385, 382, 3, 2, 2, 2, 385, 383,
	3, 2, 2, 2, 385, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 388, 7, 52,
	2, 2, 388, 89, 3, 2, 2, 2, 389, 390, 7, 4, 2, 2, 390, 91, 3, 2, 2, 2, 391,
	392, 7, 5, 2, 2, 392, 93, 3, 2, 2, 2, 393, 394, 7, 6, 2, 2, 394, 95, 3,
	2, 2, 2, 395, 396, 7, 7, 2, 2, 396, 97, 3, 2, 2, 2, 34, 101, 106, 109,
	127, 131, 144, 152, 154, 162, 166, 172, 182, 184, 193, 203, 205, 214, 218,
	223, 227, 247, 251, 295, 304, 314, 318, 322, 333, 344, 351, 358, 385,
}
var literalNames = []string{
	"", "','", "'@name'", "'@id'", "'@desc'", "'@sal'", "", "", "", "", "",
	"", "", "", "'&&'", "'||'", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "'+'", "'-'", "'/'", "'*'", "'=='", "'>'", "'<'", "'>='",
	"'<='", "'!='", "'!'", "':='", "'='", "'+='", "'-='", "'*='", "'/='", "'['",
	"']'", "';'", "'{'", "'}'", "'('", "')'", "'.'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "COUNT", "MINTE", "DAY", "SECOND", "WITHIN", "REGX",
	"NIL", "RULE", "AND", "OR", "CONC", "IF", "ELSE", "RETURN", "FOR", "BREAK",
	"FORRANGE", "CONTINUE", "TRUE", "FALSE", "NULL_LITERAL", "SALIENCE", "BEGIN",
	"END", "SIMPLENAME", "INT", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", "GT",
	"LT", "GTE", "LTE", "NOTEQUALS", "NOT", "ASSIGN", "SET", "PLUSEQUAL", "MINUSEQUAL",
	"MULTIEQUAL", "DIVEQUAL", "LSQARE", "RSQARE", "SEMICOLON", "LR_BRACE",
	"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", "DOTTEDNAME",
	"DOUBLEDOTTEDNAME", "REAL_LITERAL", "SL_COMMENT", "WS",
}

var ruleNames = []string{
	"primary", "ruleEntity", "ruleName", "ruleDescription", "salience", "ruleContent",
	"statements", "statement", "concStatement", "expression", "mathExpression",
	"expressionAtom", "assignment", "returnStmt", "aggStmt", "ifStmt", "elseIfStmt",
	"elseStmt", "forStmt", "breakStmt", "forRangeStmt", "continueStmt", "constant",
	"functionArgs", "integer", "timeForm", "expireTime", "extreValue", "realLiteral",
	"stringLiteral", "booleanLiteral", "functionCall", "methodCall", "threeLevelCall",
	"variable", "mathPmOperator", "mathMdOperator", "comparisonOperator", "countOperator",
	"logicalOperator", "assignOperator", "rangeOperator", "notOperator", "mapVar",
	"atName", "atId", "atDesc", "atSal",
}

type gengineParser struct {
	*antlr.BaseParser
}

// NewgengineParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *gengineParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewgengineParser(input antlr.TokenStream) *gengineParser {
	this := new(gengineParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "gengine.g4"

	return this
}

// gengineParser tokens.
const (
	gengineParserEOF              = antlr.TokenEOF
	gengineParserT__0             = 1
	gengineParserT__1             = 2
	gengineParserT__2             = 3
	gengineParserT__3             = 4
	gengineParserT__4             = 5
	gengineParserCOUNT            = 6
	gengineParserMINTE            = 7
	gengineParserDAY              = 8
	gengineParserSECOND           = 9
	gengineParserWITHIN           = 10
	gengineParserREGX             = 11
	gengineParserNIL              = 12
	gengineParserRULE             = 13
	gengineParserAND              = 14
	gengineParserOR               = 15
	gengineParserCONC             = 16
	gengineParserIF               = 17
	gengineParserELSE             = 18
	gengineParserRETURN           = 19
	gengineParserFOR              = 20
	gengineParserBREAK            = 21
	gengineParserFORRANGE         = 22
	gengineParserCONTINUE         = 23
	gengineParserTRUE             = 24
	gengineParserFALSE            = 25
	gengineParserNULL_LITERAL     = 26
	gengineParserSALIENCE         = 27
	gengineParserBEGIN            = 28
	gengineParserEND              = 29
	gengineParserSIMPLENAME       = 30
	gengineParserINT              = 31
	gengineParserPLUS             = 32
	gengineParserMINUS            = 33
	gengineParserDIV              = 34
	gengineParserMUL              = 35
	gengineParserEQUALS           = 36
	gengineParserGT               = 37
	gengineParserLT               = 38
	gengineParserGTE              = 39
	gengineParserLTE              = 40
	gengineParserNOTEQUALS        = 41
	gengineParserNOT              = 42
	gengineParserASSIGN           = 43
	gengineParserSET              = 44
	gengineParserPLUSEQUAL        = 45
	gengineParserMINUSEQUAL       = 46
	gengineParserMULTIEQUAL       = 47
	gengineParserDIVEQUAL         = 48
	gengineParserLSQARE           = 49
	gengineParserRSQARE           = 50
	gengineParserSEMICOLON        = 51
	gengineParserLR_BRACE         = 52
	gengineParserRR_BRACE         = 53
	gengineParserLR_BRACKET       = 54
	gengineParserRR_BRACKET       = 55
	gengineParserDOT              = 56
	gengineParserDQUOTA_STRING    = 57
	gengineParserDOTTEDNAME       = 58
	gengineParserDOUBLEDOTTEDNAME = 59
	gengineParserREAL_LITERAL     = 60
	gengineParserSL_COMMENT       = 61
	gengineParserWS               = 62
)

// gengineParser rules.
const (
	gengineParserRULE_primary            = 0
	gengineParserRULE_ruleEntity         = 1
	gengineParserRULE_ruleName           = 2
	gengineParserRULE_ruleDescription    = 3
	gengineParserRULE_salience           = 4
	gengineParserRULE_ruleContent        = 5
	gengineParserRULE_statements         = 6
	gengineParserRULE_statement          = 7
	gengineParserRULE_concStatement      = 8
	gengineParserRULE_expression         = 9
	gengineParserRULE_mathExpression     = 10
	gengineParserRULE_expressionAtom     = 11
	gengineParserRULE_assignment         = 12
	gengineParserRULE_returnStmt         = 13
	gengineParserRULE_aggStmt            = 14
	gengineParserRULE_ifStmt             = 15
	gengineParserRULE_elseIfStmt         = 16
	gengineParserRULE_elseStmt           = 17
	gengineParserRULE_forStmt            = 18
	gengineParserRULE_breakStmt          = 19
	gengineParserRULE_forRangeStmt       = 20
	gengineParserRULE_continueStmt       = 21
	gengineParserRULE_constant           = 22
	gengineParserRULE_functionArgs       = 23
	gengineParserRULE_integer            = 24
	gengineParserRULE_timeForm           = 25
	gengineParserRULE_expireTime         = 26
	gengineParserRULE_extreValue         = 27
	gengineParserRULE_realLiteral        = 28
	gengineParserRULE_stringLiteral      = 29
	gengineParserRULE_booleanLiteral     = 30
	gengineParserRULE_functionCall       = 31
	gengineParserRULE_methodCall         = 32
	gengineParserRULE_threeLevelCall     = 33
	gengineParserRULE_variable           = 34
	gengineParserRULE_mathPmOperator     = 35
	gengineParserRULE_mathMdOperator     = 36
	gengineParserRULE_comparisonOperator = 37
	gengineParserRULE_countOperator      = 38
	gengineParserRULE_logicalOperator    = 39
	gengineParserRULE_assignOperator     = 40
	gengineParserRULE_rangeOperator      = 41
	gengineParserRULE_notOperator        = 42
	gengineParserRULE_mapVar             = 43
	gengineParserRULE_atName             = 44
	gengineParserRULE_atId               = 45
	gengineParserRULE_atDesc             = 46
	gengineParserRULE_atSal              = 47
)

// IPrimaryContext is an interface to support dynamic dispatch.
type IPrimaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrimaryContext differentiates from other interfaces.
	IsPrimaryContext()
}

type PrimaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryContext() *PrimaryContext {
	var p = new(PrimaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_primary
	return p
}

func (*PrimaryContext) IsPrimaryContext() {}

func NewPrimaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryContext {
	var p = new(PrimaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_primary

	return p
}

func (s *PrimaryContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryContext) AllRuleEntity() []IRuleEntityContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRuleEntityContext)(nil)).Elem())
	var tst = make([]IRuleEntityContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRuleEntityContext)
		}
	}

	return tst
}

func (s *PrimaryContext) RuleEntity(i int) IRuleEntityContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleEntityContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRuleEntityContext)
}

func (s *PrimaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterPrimary(s)
	}
}

func (s *PrimaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitPrimary(s)
	}
}

func (s *PrimaryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitPrimary(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Primary() (localctx IPrimaryContext) {
	localctx = NewPrimaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, gengineParserRULE_primary)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == gengineParserRULE {
		{
			p.SetState(96)
			p.RuleEntity()
		}

		p.SetState(99)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IRuleEntityContext is an interface to support dynamic dispatch.
type IRuleEntityContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleEntityContext differentiates from other interfaces.
	IsRuleEntityContext()
}

type RuleEntityContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleEntityContext() *RuleEntityContext {
	var p = new(RuleEntityContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleEntity
	return p
}

func (*RuleEntityContext) IsRuleEntityContext() {}

func NewRuleEntityContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleEntityContext {
	var p = new(RuleEntityContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleEntity

	return p
}

func (s *RuleEntityContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleEntityContext) RULE() antlr.TerminalNode {
	return s.GetToken(gengineParserRULE, 0)
}

func (s *RuleEntityContext) RuleName() IRuleNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleNameContext)
}

func (s *RuleEntityContext) BEGIN() antlr.TerminalNode {
	return s.GetToken(gengineParserBEGIN, 0)
}

func (s *RuleEntityContext) RuleContent() IRuleContentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleContentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleContentContext)
}

func (s *RuleEntityContext) END() antlr.TerminalNode {
	return s.GetToken(gengineParserEND, 0)
}

func (s *RuleEntityContext) RuleDescription() IRuleDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRuleDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRuleDescriptionContext)
}

func (s *RuleEntityContext) Salience() ISalienceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISalienceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISalienceContext)
}

func (s *RuleEntityContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleEntityContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleEntityContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleEntity(s)
	}
}

func (s *RuleEntityContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleEntity(s)
	}
}

func (s *RuleEntityContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleEntity(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleEntity() (localctx IRuleEntityContext) {
	localctx = NewRuleEntityContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, gengineParserRULE_ruleEntity)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		p.Match(gengineParserRULE)
	}
	{
		p.SetState(102)
		p.RuleName()
	}
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserDQUOTA_STRING {
		{
			p.SetState(103)
			p.RuleDescription()
		}

	}
	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserSALIENCE {
		{
			p.SetState(106)
			p.Salience()
		}

	}
	{
		p.SetState(109)
		p.Match(gengineParserBEGIN)
	}
	{
		p.SetState(110)
		p.RuleContent()
	}
	{
		p.SetState(111)
		p.Match(gengineParserEND)
	}

	return localctx
}

// IRuleNameContext is an interface to support dynamic dispatch.
type IRuleNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleNameContext differentiates from other interfaces.
	IsRuleNameContext()
}

type RuleNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleNameContext() *RuleNameContext {
	var p = new(RuleNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleName
	return p
}

func (*RuleNameContext) IsRuleNameContext() {}

func NewRuleNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleNameContext {
	var p = new(RuleNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleName

	return p
}

func (s *RuleNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleNameContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *RuleNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleName(s)
	}
}

func (s *RuleNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleName(s)
	}
}

func (s *RuleNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleName() (localctx IRuleNameContext) {
	localctx = NewRuleNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, gengineParserRULE_ruleName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(113)
		p.StringLiteral()
	}

	return localctx
}

// IRuleDescriptionContext is an interface to support dynamic dispatch.
type IRuleDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleDescriptionContext differentiates from other interfaces.
	IsRuleDescriptionContext()
}

type RuleDescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleDescriptionContext() *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleDescription
	return p
}

func (*RuleDescriptionContext) IsRuleDescriptionContext() {}

func NewRuleDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleDescriptionContext {
	var p = new(RuleDescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleDescription

	return p
}

func (s *RuleDescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleDescriptionContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *RuleDescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleDescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleDescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleDescription(s)
	}
}

func (s *RuleDescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleDescription() (localctx IRuleDescriptionContext) {
	localctx = NewRuleDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, gengineParserRULE_ruleDescription)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.StringLiteral()
	}

	return localctx
}

// ISalienceContext is an interface to support dynamic dispatch.
type ISalienceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSalienceContext differentiates from other interfaces.
	IsSalienceContext()
}

type SalienceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySalienceContext() *SalienceContext {
	var p = new(SalienceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_salience
	return p
}

func (*SalienceContext) IsSalienceContext() {}

func NewSalienceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SalienceContext {
	var p = new(SalienceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_salience

	return p
}

func (s *SalienceContext) GetParser() antlr.Parser { return s.parser }

func (s *SalienceContext) SALIENCE() antlr.TerminalNode {
	return s.GetToken(gengineParserSALIENCE, 0)
}

func (s *SalienceContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *SalienceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SalienceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SalienceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterSalience(s)
	}
}

func (s *SalienceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitSalience(s)
	}
}

func (s *SalienceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitSalience(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Salience() (localctx ISalienceContext) {
	localctx = NewSalienceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, gengineParserRULE_salience)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		p.Match(gengineParserSALIENCE)
	}
	{
		p.SetState(118)
		p.Integer()
	}

	return localctx
}

// IRuleContentContext is an interface to support dynamic dispatch.
type IRuleContentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRuleContentContext differentiates from other interfaces.
	IsRuleContentContext()
}

type RuleContentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRuleContentContext() *RuleContentContext {
	var p = new(RuleContentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ruleContent
	return p
}

func (*RuleContentContext) IsRuleContentContext() {}

func NewRuleContentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RuleContentContext {
	var p = new(RuleContentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ruleContent

	return p
}

func (s *RuleContentContext) GetParser() antlr.Parser { return s.parser }

func (s *RuleContentContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *RuleContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RuleContentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RuleContentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRuleContent(s)
	}
}

func (s *RuleContentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRuleContent(s)
	}
}

func (s *RuleContentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRuleContent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RuleContent() (localctx IRuleContentContext) {
	localctx = NewRuleContentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, gengineParserRULE_ruleContent)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(120)
		p.Statements()
	}

	return localctx
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_statements
	return p
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) ReturnStmt() IReturnStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnStmtContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, gengineParserRULE_statements)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserCOUNT)|(1<<gengineParserCONC)|(1<<gengineParserIF)|(1<<gengineParserFOR)|(1<<gengineParserBREAK)|(1<<gengineParserFORRANGE)|(1<<gengineParserCONTINUE)|(1<<gengineParserSIMPLENAME))) != 0) || _la == gengineParserDOTTEDNAME || _la == gengineParserDOUBLEDOTTEDNAME {
		{
			p.SetState(122)
			p.Statement()
		}

		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserRETURN {
		{
			p.SetState(128)
			p.ReturnStmt()
		}

	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) IfStmt() IIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *StatementContext) ThreeLevelCall() IThreeLevelCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *StatementContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *StatementContext) ConcStatement() IConcStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConcStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConcStatementContext)
}

func (s *StatementContext) ForStmt() IForStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForStmtContext)
}

func (s *StatementContext) BreakStmt() IBreakStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreakStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *StatementContext) ForRangeStmt() IForRangeStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IForRangeStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IForRangeStmtContext)
}

func (s *StatementContext) ContinueStmt() IContinueStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinueStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *StatementContext) AggStmt() IAggStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAggStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAggStmtContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, gengineParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(131)
			p.IfStmt()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.FunctionCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.MethodCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(134)
			p.ThreeLevelCall()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(135)
			p.Assignment()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(136)
			p.ConcStatement()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(137)
			p.ForStmt()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(138)
			p.BreakStmt()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(139)
			p.ForRangeStmt()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(140)
			p.ContinueStmt()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(141)
			p.AggStmt()
		}

	}

	return localctx
}

// IConcStatementContext is an interface to support dynamic dispatch.
type IConcStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConcStatementContext differentiates from other interfaces.
	IsConcStatementContext()
}

type ConcStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConcStatementContext() *ConcStatementContext {
	var p = new(ConcStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_concStatement
	return p
}

func (*ConcStatementContext) IsConcStatementContext() {}

func NewConcStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConcStatementContext {
	var p = new(ConcStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_concStatement

	return p
}

func (s *ConcStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ConcStatementContext) CONC() antlr.TerminalNode {
	return s.GetToken(gengineParserCONC, 0)
}

func (s *ConcStatementContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ConcStatementContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ConcStatementContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ConcStatementContext) AllMethodCall() []IMethodCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodCallContext)(nil)).Elem())
	var tst = make([]IMethodCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodCallContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) MethodCall(i int) IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ConcStatementContext) AllThreeLevelCall() []IThreeLevelCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem())
	var tst = make([]IThreeLevelCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IThreeLevelCallContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) ThreeLevelCall(i int) IThreeLevelCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *ConcStatementContext) AllAssignment() []IAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignmentContext)(nil)).Elem())
	var tst = make([]IAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignmentContext)
		}
	}

	return tst
}

func (s *ConcStatementContext) Assignment(i int) IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ConcStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConcStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConcStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterConcStatement(s)
	}
}

func (s *ConcStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitConcStatement(s)
	}
}

func (s *ConcStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitConcStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ConcStatement() (localctx IConcStatementContext) {
	localctx = NewConcStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, gengineParserRULE_concStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(gengineParserCONC)
	}
	{
		p.SetState(145)
		p.Match(gengineParserLR_BRACE)
	}
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(gengineParserSIMPLENAME-30))|(1<<(gengineParserDOTTEDNAME-30))|(1<<(gengineParserDOUBLEDOTTEDNAME-30)))) != 0 {
		p.SetState(150)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(146)
				p.FunctionCall()
			}

		case 2:
			{
				p.SetState(147)
				p.MethodCall()
			}

		case 3:
			{
				p.SetState(148)
				p.ThreeLevelCall()
			}

		case 4:
			{
				p.SetState(149)
				p.Assignment()
			}

		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(155)
		p.Match(gengineParserRR_BRACE)
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) MathExpression() IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *ExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionContext) NotOperator() INotOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INotOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INotOperatorContext)
}

func (s *ExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *ExpressionContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *ExpressionContext) LogicalOperator() ILogicalOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *gengineParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 18
	p.EnterRecursionRule(localctx, 18, gengineParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(170)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(158)
			p.mathExpression(0)
		}

	case 2:
		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == gengineParserNOT {
			{
				p.SetState(159)
				p.NotOperator()
			}

		}
		{
			p.SetState(162)
			p.ExpressionAtom()
		}

	case 3:
		p.SetState(164)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == gengineParserNOT {
			{
				p.SetState(163)
				p.NotOperator()
			}

		}
		{
			p.SetState(166)
			p.Match(gengineParserLR_BRACKET)
		}
		{
			p.SetState(167)
			p.expression(0)
		}
		{
			p.SetState(168)
			p.Match(gengineParserRR_BRACKET)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(180)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(172)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(173)
					p.ComparisonOperator()
				}
				{
					p.SetState(174)
					p.expression(5)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_expression)
				p.SetState(176)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(177)
					p.LogicalOperator()
				}
				{
					p.SetState(178)
					p.expression(4)
				}

			}

		}
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
	}

	return localctx
}

// IMathExpressionContext is an interface to support dynamic dispatch.
type IMathExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathExpressionContext differentiates from other interfaces.
	IsMathExpressionContext()
}

type MathExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathExpressionContext() *MathExpressionContext {
	var p = new(MathExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mathExpression
	return p
}

func (*MathExpressionContext) IsMathExpressionContext() {}

func NewMathExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathExpressionContext {
	var p = new(MathExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathExpression

	return p
}

func (s *MathExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MathExpressionContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *MathExpressionContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *MathExpressionContext) AllMathExpression() []IMathExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem())
	var tst = make([]IMathExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMathExpressionContext)
		}
	}

	return tst
}

func (s *MathExpressionContext) MathExpression(i int) IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *MathExpressionContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *MathExpressionContext) MathMdOperator() IMathMdOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathMdOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathMdOperatorContext)
}

func (s *MathExpressionContext) MathPmOperator() IMathPmOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathPmOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathPmOperatorContext)
}

func (s *MathExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathExpression(s)
	}
}

func (s *MathExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathExpression(s)
	}
}

func (s *MathExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MathExpression() (localctx IMathExpressionContext) {
	return p.mathExpression(0)
}

func (p *gengineParser) mathExpression(_p int) (localctx IMathExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewMathExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IMathExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, gengineParserRULE_mathExpression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(191)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gengineParserT__1, gengineParserT__2, gengineParserT__3, gengineParserT__4, gengineParserTRUE, gengineParserFALSE, gengineParserSIMPLENAME, gengineParserINT, gengineParserMINUS, gengineParserDQUOTA_STRING, gengineParserDOTTEDNAME, gengineParserDOUBLEDOTTEDNAME, gengineParserREAL_LITERAL:
		{
			p.SetState(186)
			p.ExpressionAtom()
		}

	case gengineParserLR_BRACKET:
		{
			p.SetState(187)
			p.Match(gengineParserLR_BRACKET)
		}
		{
			p.SetState(188)
			p.mathExpression(0)
		}
		{
			p.SetState(189)
			p.Match(gengineParserRR_BRACKET)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(201)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(193)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(194)
					p.MathMdOperator()
				}
				{
					p.SetState(195)
					p.mathExpression(5)
				}

			case 2:
				localctx = NewMathExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, gengineParserRULE_mathExpression)
				p.SetState(197)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(198)
					p.MathPmOperator()
				}
				{
					p.SetState(199)
					p.mathExpression(4)
				}

			}

		}
		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionAtomContext) MethodCall() IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *ExpressionAtomContext) ThreeLevelCall() IThreeLevelCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *ExpressionAtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ExpressionAtomContext) MapVar() IMapVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *ExpressionAtomContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExpressionAtom(s)
	}
}

func (s *ExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, gengineParserRULE_expressionAtom)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(212)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(206)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(207)
			p.MethodCall()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(208)
			p.ThreeLevelCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(209)
			p.Constant()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(210)
			p.MapVar()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(211)
			p.Variable()
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) AssignOperator() IAssignOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignOperatorContext)
}

func (s *AssignmentContext) MapVar() IMapVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapVarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *AssignmentContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *AssignmentContext) MathExpression() IMathExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathExpressionContext)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, gengineParserRULE_assignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(214)
			p.MapVar()
		}

	case 2:
		{
			p.SetState(215)
			p.Variable()
		}

	}
	{
		p.SetState(218)
		p.AssignOperator()
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(219)
			p.mathExpression(0)
		}

	case 2:
		{
			p.SetState(220)
			p.expression(0)
		}

	}

	return localctx
}

// IReturnStmtContext is an interface to support dynamic dispatch.
type IReturnStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReturnStmtContext differentiates from other interfaces.
	IsReturnStmtContext()
}

type ReturnStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStmtContext() *ReturnStmtContext {
	var p = new(ReturnStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_returnStmt
	return p
}

func (*ReturnStmtContext) IsReturnStmtContext() {}

func NewReturnStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStmtContext {
	var p = new(ReturnStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_returnStmt

	return p
}

func (s *ReturnStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStmtContext) RETURN() antlr.TerminalNode {
	return s.GetToken(gengineParserRETURN, 0)
}

func (s *ReturnStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterReturnStmt(s)
	}
}

func (s *ReturnStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitReturnStmt(s)
	}
}

func (s *ReturnStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitReturnStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ReturnStmt() (localctx IReturnStmtContext) {
	localctx = NewReturnStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, gengineParserRULE_returnStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(223)
		p.Match(gengineParserRETURN)
	}
	p.SetState(225)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__1)|(1<<gengineParserT__2)|(1<<gengineParserT__3)|(1<<gengineParserT__4)|(1<<gengineParserTRUE)|(1<<gengineParserFALSE)|(1<<gengineParserSIMPLENAME)|(1<<gengineParserINT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(gengineParserMINUS-33))|(1<<(gengineParserNOT-33))|(1<<(gengineParserLR_BRACKET-33))|(1<<(gengineParserDQUOTA_STRING-33))|(1<<(gengineParserDOTTEDNAME-33))|(1<<(gengineParserDOUBLEDOTTEDNAME-33))|(1<<(gengineParserREAL_LITERAL-33)))) != 0) {
		{
			p.SetState(224)
			p.expression(0)
		}

	}

	return localctx
}

// IAggStmtContext is an interface to support dynamic dispatch.
type IAggStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAggStmtContext differentiates from other interfaces.
	IsAggStmtContext()
}

type AggStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggStmtContext() *AggStmtContext {
	var p = new(AggStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_aggStmt
	return p
}

func (*AggStmtContext) IsAggStmtContext() {}

func NewAggStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggStmtContext {
	var p = new(AggStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_aggStmt

	return p
}

func (s *AggStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AggStmtContext) COUNT() antlr.TerminalNode {
	return s.GetToken(gengineParserCOUNT, 0)
}

func (s *AggStmtContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *AggStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AggStmtContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *AggStmtContext) CountOperator() ICountOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICountOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICountOperatorContext)
}

func (s *AggStmtContext) ExtreValue() IExtreValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtreValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtreValueContext)
}

func (s *AggStmtContext) WITHIN() antlr.TerminalNode {
	return s.GetToken(gengineParserWITHIN, 0)
}

func (s *AggStmtContext) ExpireTime() IExpireTimeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpireTimeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpireTimeContext)
}

func (s *AggStmtContext) TimeForm() ITimeFormContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITimeFormContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITimeFormContext)
}

func (s *AggStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAggStmt(s)
	}
}

func (s *AggStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAggStmt(s)
	}
}

func (s *AggStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAggStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AggStmt() (localctx IAggStmtContext) {
	localctx = NewAggStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, gengineParserRULE_aggStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(227)
		p.Match(gengineParserCOUNT)
	}
	{
		p.SetState(228)
		p.Match(gengineParserLR_BRACKET)
	}
	{
		p.SetState(229)
		p.expression(0)
	}
	{
		p.SetState(230)
		p.Match(gengineParserRR_BRACKET)
	}
	{
		p.SetState(231)
		p.CountOperator()
	}
	{
		p.SetState(232)
		p.ExtreValue()
	}
	{
		p.SetState(233)
		p.Match(gengineParserWITHIN)
	}
	{
		p.SetState(234)
		p.ExpireTime()
	}
	{
		p.SetState(235)
		p.TimeForm()
	}

	return localctx
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_ifStmt
	return p
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(gengineParserIF, 0)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *IfStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *IfStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *IfStmtContext) AllElseIfStmt() []IElseIfStmtContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElseIfStmtContext)(nil)).Elem())
	var tst = make([]IElseIfStmtContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElseIfStmtContext)
		}
	}

	return tst
}

func (s *IfStmtContext) ElseIfStmt(i int) IElseIfStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseIfStmtContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElseIfStmtContext)
}

func (s *IfStmtContext) ElseStmt() IElseStmtContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStmtContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, gengineParserRULE_ifStmt)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(237)
		p.Match(gengineParserIF)
	}
	{
		p.SetState(238)
		p.expression(0)
	}
	{
		p.SetState(239)
		p.Match(gengineParserLR_BRACE)
	}
	{
		p.SetState(240)
		p.Statements()
	}
	{
		p.SetState(241)
		p.Match(gengineParserRR_BRACE)
	}
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(242)
				p.ElseIfStmt()
			}

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext())
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserELSE {
		{
			p.SetState(248)
			p.ElseStmt()
		}

	}

	return localctx
}

// IElseIfStmtContext is an interface to support dynamic dispatch.
type IElseIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseIfStmtContext differentiates from other interfaces.
	IsElseIfStmtContext()
}

type ElseIfStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfStmtContext() *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_elseIfStmt
	return p
}

func (*ElseIfStmtContext) IsElseIfStmtContext() {}

func NewElseIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfStmtContext {
	var p = new(ElseIfStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_elseIfStmt

	return p
}

func (s *ElseIfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(gengineParserELSE, 0)
}

func (s *ElseIfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(gengineParserIF, 0)
}

func (s *ElseIfStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ElseIfStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ElseIfStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ElseIfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitElseIfStmt(s)
	}
}

func (s *ElseIfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitElseIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ElseIfStmt() (localctx IElseIfStmtContext) {
	localctx = NewElseIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, gengineParserRULE_elseIfStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(gengineParserELSE)
	}
	{
		p.SetState(252)
		p.Match(gengineParserIF)
	}
	{
		p.SetState(253)
		p.expression(0)
	}
	{
		p.SetState(254)
		p.Match(gengineParserLR_BRACE)
	}
	{
		p.SetState(255)
		p.Statements()
	}
	{
		p.SetState(256)
		p.Match(gengineParserRR_BRACE)
	}

	return localctx
}

// IElseStmtContext is an interface to support dynamic dispatch.
type IElseStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStmtContext differentiates from other interfaces.
	IsElseStmtContext()
}

type ElseStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStmtContext() *ElseStmtContext {
	var p = new(ElseStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_elseStmt
	return p
}

func (*ElseStmtContext) IsElseStmtContext() {}

func NewElseStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStmtContext {
	var p = new(ElseStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_elseStmt

	return p
}

func (s *ElseStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(gengineParserELSE, 0)
}

func (s *ElseStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ElseStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ElseStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ElseStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterElseStmt(s)
	}
}

func (s *ElseStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitElseStmt(s)
	}
}

func (s *ElseStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitElseStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ElseStmt() (localctx IElseStmtContext) {
	localctx = NewElseStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, gengineParserRULE_elseStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.Match(gengineParserELSE)
	}
	{
		p.SetState(259)
		p.Match(gengineParserLR_BRACE)
	}
	{
		p.SetState(260)
		p.Statements()
	}
	{
		p.SetState(261)
		p.Match(gengineParserRR_BRACE)
	}

	return localctx
}

// IForStmtContext is an interface to support dynamic dispatch.
type IForStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForStmtContext differentiates from other interfaces.
	IsForStmtContext()
}

type ForStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStmtContext() *ForStmtContext {
	var p = new(ForStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_forStmt
	return p
}

func (*ForStmtContext) IsForStmtContext() {}

func NewForStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStmtContext {
	var p = new(ForStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_forStmt

	return p
}

func (s *ForStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStmtContext) FOR() antlr.TerminalNode {
	return s.GetToken(gengineParserFOR, 0)
}

func (s *ForStmtContext) AllAssignment() []IAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAssignmentContext)(nil)).Elem())
	var tst = make([]IAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAssignmentContext)
		}
	}

	return tst
}

func (s *ForStmtContext) Assignment(i int) IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ForStmtContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(gengineParserSEMICOLON)
}

func (s *ForStmtContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(gengineParserSEMICOLON, i)
}

func (s *ForStmtContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ForStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ForStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ForStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterForStmt(s)
	}
}

func (s *ForStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitForStmt(s)
	}
}

func (s *ForStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitForStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ForStmt() (localctx IForStmtContext) {
	localctx = NewForStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, gengineParserRULE_forStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(gengineParserFOR)
	}
	{
		p.SetState(264)
		p.Assignment()
	}
	{
		p.SetState(265)
		p.Match(gengineParserSEMICOLON)
	}
	{
		p.SetState(266)
		p.expression(0)
	}
	{
		p.SetState(267)
		p.Match(gengineParserSEMICOLON)
	}
	{
		p.SetState(268)
		p.Assignment()
	}
	{
		p.SetState(269)
		p.Match(gengineParserLR_BRACE)
	}
	{
		p.SetState(270)
		p.Statements()
	}
	{
		p.SetState(271)
		p.Match(gengineParserRR_BRACE)
	}

	return localctx
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_breakStmt
	return p
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(gengineParserBREAK, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, gengineParserRULE_breakStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(273)
		p.Match(gengineParserBREAK)
	}

	return localctx
}

// IForRangeStmtContext is an interface to support dynamic dispatch.
type IForRangeStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsForRangeStmtContext differentiates from other interfaces.
	IsForRangeStmtContext()
}

type ForRangeStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForRangeStmtContext() *ForRangeStmtContext {
	var p = new(ForRangeStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_forRangeStmt
	return p
}

func (*ForRangeStmtContext) IsForRangeStmtContext() {}

func NewForRangeStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForRangeStmtContext {
	var p = new(ForRangeStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_forRangeStmt

	return p
}

func (s *ForRangeStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ForRangeStmtContext) FORRANGE() antlr.TerminalNode {
	return s.GetToken(gengineParserFORRANGE, 0)
}

func (s *ForRangeStmtContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *ForRangeStmtContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ForRangeStmtContext) RangeOperator() IRangeOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangeOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRangeOperatorContext)
}

func (s *ForRangeStmtContext) LR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACE, 0)
}

func (s *ForRangeStmtContext) Statements() IStatementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ForRangeStmtContext) RR_BRACE() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACE, 0)
}

func (s *ForRangeStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForRangeStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForRangeStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterForRangeStmt(s)
	}
}

func (s *ForRangeStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitForRangeStmt(s)
	}
}

func (s *ForRangeStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitForRangeStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ForRangeStmt() (localctx IForRangeStmtContext) {
	localctx = NewForRangeStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, gengineParserRULE_forRangeStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(275)
		p.Match(gengineParserFORRANGE)
	}
	{
		p.SetState(276)
		p.Variable()
	}
	{
		p.SetState(277)
		p.RangeOperator()
	}
	{
		p.SetState(278)
		p.Variable()
	}
	{
		p.SetState(279)
		p.Match(gengineParserLR_BRACE)
	}
	{
		p.SetState(280)
		p.Statements()
	}
	{
		p.SetState(281)
		p.Match(gengineParserRR_BRACE)
	}

	return localctx
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_continueStmt
	return p
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(gengineParserCONTINUE, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, gengineParserRULE_continueStmt)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(gengineParserCONTINUE)
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *ConstantContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *ConstantContext) RealLiteral() IRealLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRealLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *ConstantContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *ConstantContext) AtName() IAtNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtNameContext)
}

func (s *ConstantContext) AtId() IAtIdContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtIdContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtIdContext)
}

func (s *ConstantContext) AtDesc() IAtDescContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtDescContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtDescContext)
}

func (s *ConstantContext) AtSal() IAtSalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtSalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtSalContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitConstant(s)
	}
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, gengineParserRULE_constant)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(293)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(285)
			p.BooleanLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(286)
			p.Integer()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(287)
			p.RealLiteral()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(288)
			p.StringLiteral()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(289)
			p.AtName()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(290)
			p.AtId()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(291)
			p.AtDesc()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(292)
			p.AtSal()
		}

	}

	return localctx
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) AllConstant() []IConstantContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IConstantContext)(nil)).Elem())
	var tst = make([]IConstantContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IConstantContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Constant(i int) IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FunctionArgsContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *FunctionArgsContext) AllFunctionCall() []IFunctionCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem())
	var tst = make([]IFunctionCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) FunctionCall(i int) IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionArgsContext) AllMethodCall() []IMethodCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodCallContext)(nil)).Elem())
	var tst = make([]IMethodCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) MethodCall(i int) IMethodCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodCallContext)
}

func (s *FunctionArgsContext) AllThreeLevelCall() []IThreeLevelCallContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem())
	var tst = make([]IThreeLevelCallContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IThreeLevelCallContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) ThreeLevelCall(i int) IThreeLevelCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IThreeLevelCallContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IThreeLevelCallContext)
}

func (s *FunctionArgsContext) AllMapVar() []IMapVarContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMapVarContext)(nil)).Elem())
	var tst = make([]IMapVarContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMapVarContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) MapVar(i int) IMapVarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMapVarContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMapVarContext)
}

func (s *FunctionArgsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitFunctionArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, gengineParserRULE_functionArgs)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(295)
			p.Constant()
		}

	case 2:
		{
			p.SetState(296)
			p.Variable()
		}

	case 3:
		{
			p.SetState(297)
			p.FunctionCall()
		}

	case 4:
		{
			p.SetState(298)
			p.MethodCall()
		}

	case 5:
		{
			p.SetState(299)
			p.ThreeLevelCall()
		}

	case 6:
		{
			p.SetState(300)
			p.MapVar()
		}

	case 7:
		{
			p.SetState(301)
			p.expression(0)
		}

	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == gengineParserT__0 {
		{
			p.SetState(304)
			p.Match(gengineParserT__0)
		}
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(305)
				p.Constant()
			}

		case 2:
			{
				p.SetState(306)
				p.Variable()
			}

		case 3:
			{
				p.SetState(307)
				p.FunctionCall()
			}

		case 4:
			{
				p.SetState(308)
				p.MethodCall()
			}

		case 5:
			{
				p.SetState(309)
				p.ThreeLevelCall()
			}

		case 6:
			{
				p.SetState(310)
				p.MapVar()
			}

		case 7:
			{
				p.SetState(311)
				p.expression(0)
			}

		}

		p.SetState(318)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IIntegerContext is an interface to support dynamic dispatch.
type IIntegerContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerContext differentiates from other interfaces.
	IsIntegerContext()
}

type IntegerContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerContext() *IntegerContext {
	var p = new(IntegerContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_integer
	return p
}

func (*IntegerContext) IsIntegerContext() {}

func NewIntegerContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerContext {
	var p = new(IntegerContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_integer

	return p
}

func (s *IntegerContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerContext) INT() antlr.TerminalNode {
	return s.GetToken(gengineParserINT, 0)
}

func (s *IntegerContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *IntegerContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterInteger(s)
	}
}

func (s *IntegerContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitInteger(s)
	}
}

func (s *IntegerContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitInteger(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Integer() (localctx IIntegerContext) {
	localctx = NewIntegerContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, gengineParserRULE_integer)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserMINUS {
		{
			p.SetState(319)
			p.Match(gengineParserMINUS)
		}

	}
	{
		p.SetState(322)
		p.Match(gengineParserINT)
	}

	return localctx
}

// ITimeFormContext is an interface to support dynamic dispatch.
type ITimeFormContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTimeFormContext differentiates from other interfaces.
	IsTimeFormContext()
}

type TimeFormContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeFormContext() *TimeFormContext {
	var p = new(TimeFormContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_timeForm
	return p
}

func (*TimeFormContext) IsTimeFormContext() {}

func NewTimeFormContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeFormContext {
	var p = new(TimeFormContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_timeForm

	return p
}

func (s *TimeFormContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeFormContext) MINTE() antlr.TerminalNode {
	return s.GetToken(gengineParserMINTE, 0)
}

func (s *TimeFormContext) DAY() antlr.TerminalNode {
	return s.GetToken(gengineParserDAY, 0)
}

func (s *TimeFormContext) SECOND() antlr.TerminalNode {
	return s.GetToken(gengineParserSECOND, 0)
}

func (s *TimeFormContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeFormContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeFormContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterTimeForm(s)
	}
}

func (s *TimeFormContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitTimeForm(s)
	}
}

func (s *TimeFormContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitTimeForm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) TimeForm() (localctx ITimeFormContext) {
	localctx = NewTimeFormContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, gengineParserRULE_timeForm)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(324)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserMINTE)|(1<<gengineParserDAY)|(1<<gengineParserSECOND))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpireTimeContext is an interface to support dynamic dispatch.
type IExpireTimeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpireTimeContext differentiates from other interfaces.
	IsExpireTimeContext()
}

type ExpireTimeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpireTimeContext() *ExpireTimeContext {
	var p = new(ExpireTimeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_expireTime
	return p
}

func (*ExpireTimeContext) IsExpireTimeContext() {}

func NewExpireTimeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpireTimeContext {
	var p = new(ExpireTimeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_expireTime

	return p
}

func (s *ExpireTimeContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpireTimeContext) INT() antlr.TerminalNode {
	return s.GetToken(gengineParserINT, 0)
}

func (s *ExpireTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpireTimeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpireTimeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExpireTime(s)
	}
}

func (s *ExpireTimeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExpireTime(s)
	}
}

func (s *ExpireTimeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExpireTime(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ExpireTime() (localctx IExpireTimeContext) {
	localctx = NewExpireTimeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, gengineParserRULE_expireTime)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(326)
		p.Match(gengineParserINT)
	}

	return localctx
}

// IExtreValueContext is an interface to support dynamic dispatch.
type IExtreValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExtreValueContext differentiates from other interfaces.
	IsExtreValueContext()
}

type ExtreValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtreValueContext() *ExtreValueContext {
	var p = new(ExtreValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_extreValue
	return p
}

func (*ExtreValueContext) IsExtreValueContext() {}

func NewExtreValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtreValueContext {
	var p = new(ExtreValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_extreValue

	return p
}

func (s *ExtreValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtreValueContext) INT() antlr.TerminalNode {
	return s.GetToken(gengineParserINT, 0)
}

func (s *ExtreValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtreValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtreValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterExtreValue(s)
	}
}

func (s *ExtreValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitExtreValue(s)
	}
}

func (s *ExtreValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitExtreValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ExtreValue() (localctx IExtreValueContext) {
	localctx = NewExtreValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, gengineParserRULE_extreValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)
		p.Match(gengineParserINT)
	}

	return localctx
}

// IRealLiteralContext is an interface to support dynamic dispatch.
type IRealLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealLiteralContext differentiates from other interfaces.
	IsRealLiteralContext()
}

type RealLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealLiteralContext() *RealLiteralContext {
	var p = new(RealLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_realLiteral
	return p
}

func (*RealLiteralContext) IsRealLiteralContext() {}

func NewRealLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealLiteralContext {
	var p = new(RealLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_realLiteral

	return p
}

func (s *RealLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RealLiteralContext) REAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(gengineParserREAL_LITERAL, 0)
}

func (s *RealLiteralContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *RealLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRealLiteral(s)
	}
}

func (s *RealLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRealLiteral(s)
	}
}

func (s *RealLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRealLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RealLiteral() (localctx IRealLiteralContext) {
	localctx = NewRealLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, gengineParserRULE_realLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == gengineParserMINUS {
		{
			p.SetState(330)
			p.Match(gengineParserMINUS)
		}

	}
	{
		p.SetState(333)
		p.Match(gengineParserREAL_LITERAL)
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) DQUOTA_STRING() antlr.TerminalNode {
	return s.GetToken(gengineParserDQUOTA_STRING, 0)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, gengineParserRULE_stringLiteral)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(gengineParserDQUOTA_STRING)
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(gengineParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(gengineParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, gengineParserRULE_booleanLiteral)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(337)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserTRUE || _la == gengineParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(gengineParserSIMPLENAME, 0)
}

func (s *FunctionCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *FunctionCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *FunctionCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, gengineParserRULE_functionCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(gengineParserSIMPLENAME)
	}
	{
		p.SetState(340)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__1)|(1<<gengineParserT__2)|(1<<gengineParserT__3)|(1<<gengineParserT__4)|(1<<gengineParserTRUE)|(1<<gengineParserFALSE)|(1<<gengineParserSIMPLENAME)|(1<<gengineParserINT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(gengineParserMINUS-33))|(1<<(gengineParserNOT-33))|(1<<(gengineParserLR_BRACKET-33))|(1<<(gengineParserDQUOTA_STRING-33))|(1<<(gengineParserDOTTEDNAME-33))|(1<<(gengineParserDOUBLEDOTTEDNAME-33))|(1<<(gengineParserREAL_LITERAL-33)))) != 0) {
		{
			p.SetState(341)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(344)
		p.Match(gengineParserRR_BRACKET)
	}

	return localctx
}

// IMethodCallContext is an interface to support dynamic dispatch.
type IMethodCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMethodCallContext differentiates from other interfaces.
	IsMethodCallContext()
}

type MethodCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMethodCallContext() *MethodCallContext {
	var p = new(MethodCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_methodCall
	return p
}

func (*MethodCallContext) IsMethodCallContext() {}

func NewMethodCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodCallContext {
	var p = new(MethodCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_methodCall

	return p
}

func (s *MethodCallContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodCallContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOTTEDNAME, 0)
}

func (s *MethodCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *MethodCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *MethodCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *MethodCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMethodCall(s)
	}
}

func (s *MethodCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMethodCall(s)
	}
}

func (s *MethodCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMethodCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MethodCall() (localctx IMethodCallContext) {
	localctx = NewMethodCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, gengineParserRULE_methodCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.Match(gengineParserDOTTEDNAME)
	}
	{
		p.SetState(347)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__1)|(1<<gengineParserT__2)|(1<<gengineParserT__3)|(1<<gengineParserT__4)|(1<<gengineParserTRUE)|(1<<gengineParserFALSE)|(1<<gengineParserSIMPLENAME)|(1<<gengineParserINT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(gengineParserMINUS-33))|(1<<(gengineParserNOT-33))|(1<<(gengineParserLR_BRACKET-33))|(1<<(gengineParserDQUOTA_STRING-33))|(1<<(gengineParserDOTTEDNAME-33))|(1<<(gengineParserDOUBLEDOTTEDNAME-33))|(1<<(gengineParserREAL_LITERAL-33)))) != 0) {
		{
			p.SetState(348)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(351)
		p.Match(gengineParserRR_BRACKET)
	}

	return localctx
}

// IThreeLevelCallContext is an interface to support dynamic dispatch.
type IThreeLevelCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsThreeLevelCallContext differentiates from other interfaces.
	IsThreeLevelCallContext()
}

type ThreeLevelCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyThreeLevelCallContext() *ThreeLevelCallContext {
	var p = new(ThreeLevelCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_threeLevelCall
	return p
}

func (*ThreeLevelCallContext) IsThreeLevelCallContext() {}

func NewThreeLevelCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ThreeLevelCallContext {
	var p = new(ThreeLevelCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_threeLevelCall

	return p
}

func (s *ThreeLevelCallContext) GetParser() antlr.Parser { return s.parser }

func (s *ThreeLevelCallContext) DOUBLEDOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOUBLEDOTTEDNAME, 0)
}

func (s *ThreeLevelCallContext) LR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserLR_BRACKET, 0)
}

func (s *ThreeLevelCallContext) RR_BRACKET() antlr.TerminalNode {
	return s.GetToken(gengineParserRR_BRACKET, 0)
}

func (s *ThreeLevelCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *ThreeLevelCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ThreeLevelCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ThreeLevelCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterThreeLevelCall(s)
	}
}

func (s *ThreeLevelCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitThreeLevelCall(s)
	}
}

func (s *ThreeLevelCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitThreeLevelCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ThreeLevelCall() (localctx IThreeLevelCallContext) {
	localctx = NewThreeLevelCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, gengineParserRULE_threeLevelCall)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(353)
		p.Match(gengineParserDOUBLEDOTTEDNAME)
	}
	{
		p.SetState(354)
		p.Match(gengineParserLR_BRACKET)
	}
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<gengineParserT__1)|(1<<gengineParserT__2)|(1<<gengineParserT__3)|(1<<gengineParserT__4)|(1<<gengineParserTRUE)|(1<<gengineParserFALSE)|(1<<gengineParserSIMPLENAME)|(1<<gengineParserINT))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(gengineParserMINUS-33))|(1<<(gengineParserNOT-33))|(1<<(gengineParserLR_BRACKET-33))|(1<<(gengineParserDQUOTA_STRING-33))|(1<<(gengineParserDOTTEDNAME-33))|(1<<(gengineParserDOUBLEDOTTEDNAME-33))|(1<<(gengineParserREAL_LITERAL-33)))) != 0) {
		{
			p.SetState(355)
			p.FunctionArgs()
		}

	}
	{
		p.SetState(358)
		p.Match(gengineParserRR_BRACKET)
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) SIMPLENAME() antlr.TerminalNode {
	return s.GetToken(gengineParserSIMPLENAME, 0)
}

func (s *VariableContext) DOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOTTEDNAME, 0)
}

func (s *VariableContext) DOUBLEDOTTEDNAME() antlr.TerminalNode {
	return s.GetToken(gengineParserDOUBLEDOTTEDNAME, 0)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, gengineParserRULE_variable)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(360)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(gengineParserSIMPLENAME-30))|(1<<(gengineParserDOTTEDNAME-30))|(1<<(gengineParserDOUBLEDOTTEDNAME-30)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMathPmOperatorContext is an interface to support dynamic dispatch.
type IMathPmOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathPmOperatorContext differentiates from other interfaces.
	IsMathPmOperatorContext()
}

type MathPmOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathPmOperatorContext() *MathPmOperatorContext {
	var p = new(MathPmOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mathPmOperator
	return p
}

func (*MathPmOperatorContext) IsMathPmOperatorContext() {}

func NewMathPmOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathPmOperatorContext {
	var p = new(MathPmOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathPmOperator

	return p
}

func (s *MathPmOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathPmOperatorContext) PLUS() antlr.TerminalNode {
	return s.GetToken(gengineParserPLUS, 0)
}

func (s *MathPmOperatorContext) MINUS() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUS, 0)
}

func (s *MathPmOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathPmOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathPmOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathPmOperator(s)
	}
}

func (s *MathPmOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathPmOperator(s)
	}
}

func (s *MathPmOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathPmOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MathPmOperator() (localctx IMathPmOperatorContext) {
	localctx = NewMathPmOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, gengineParserRULE_mathPmOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserPLUS || _la == gengineParserMINUS) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMathMdOperatorContext is an interface to support dynamic dispatch.
type IMathMdOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathMdOperatorContext differentiates from other interfaces.
	IsMathMdOperatorContext()
}

type MathMdOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathMdOperatorContext() *MathMdOperatorContext {
	var p = new(MathMdOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mathMdOperator
	return p
}

func (*MathMdOperatorContext) IsMathMdOperatorContext() {}

func NewMathMdOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathMdOperatorContext {
	var p = new(MathMdOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mathMdOperator

	return p
}

func (s *MathMdOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *MathMdOperatorContext) MUL() antlr.TerminalNode {
	return s.GetToken(gengineParserMUL, 0)
}

func (s *MathMdOperatorContext) DIV() antlr.TerminalNode {
	return s.GetToken(gengineParserDIV, 0)
}

func (s *MathMdOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathMdOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathMdOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMathMdOperator(s)
	}
}

func (s *MathMdOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMathMdOperator(s)
	}
}

func (s *MathMdOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMathMdOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MathMdOperator() (localctx IMathMdOperatorContext) {
	localctx = NewMathMdOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, gengineParserRULE_mathMdOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(364)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserDIV || _la == gengineParserMUL) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(gengineParserGT, 0)
}

func (s *ComparisonOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(gengineParserLT, 0)
}

func (s *ComparisonOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(gengineParserGTE, 0)
}

func (s *ComparisonOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(gengineParserLTE, 0)
}

func (s *ComparisonOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserEQUALS, 0)
}

func (s *ComparisonOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserNOTEQUALS, 0)
}

func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, gengineParserRULE_comparisonOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(gengineParserEQUALS-36))|(1<<(gengineParserGT-36))|(1<<(gengineParserLT-36))|(1<<(gengineParserGTE-36))|(1<<(gengineParserLTE-36))|(1<<(gengineParserNOTEQUALS-36)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICountOperatorContext is an interface to support dynamic dispatch.
type ICountOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCountOperatorContext differentiates from other interfaces.
	IsCountOperatorContext()
}

type CountOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCountOperatorContext() *CountOperatorContext {
	var p = new(CountOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_countOperator
	return p
}

func (*CountOperatorContext) IsCountOperatorContext() {}

func NewCountOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CountOperatorContext {
	var p = new(CountOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_countOperator

	return p
}

func (s *CountOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *CountOperatorContext) GT() antlr.TerminalNode {
	return s.GetToken(gengineParserGT, 0)
}

func (s *CountOperatorContext) LT() antlr.TerminalNode {
	return s.GetToken(gengineParserLT, 0)
}

func (s *CountOperatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(gengineParserGTE, 0)
}

func (s *CountOperatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(gengineParserLTE, 0)
}

func (s *CountOperatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserEQUALS, 0)
}

func (s *CountOperatorContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(gengineParserNOTEQUALS, 0)
}

func (s *CountOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CountOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterCountOperator(s)
	}
}

func (s *CountOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitCountOperator(s)
	}
}

func (s *CountOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitCountOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) CountOperator() (localctx ICountOperatorContext) {
	localctx = NewCountOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, gengineParserRULE_countOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(368)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-36)&-(0x1f+1)) == 0 && ((1<<uint((_la-36)))&((1<<(gengineParserEQUALS-36))|(1<<(gengineParserGT-36))|(1<<(gengineParserLT-36))|(1<<(gengineParserGTE-36))|(1<<(gengineParserLTE-36))|(1<<(gengineParserNOTEQUALS-36)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(gengineParserAND, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(gengineParserOR, 0)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitLogicalOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, gengineParserRULE_logicalOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(370)
		_la = p.GetTokenStream().LA(1)

		if !(_la == gengineParserAND || _la == gengineParserOR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAssignOperatorContext is an interface to support dynamic dispatch.
type IAssignOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAssignOperatorContext differentiates from other interfaces.
	IsAssignOperatorContext()
}

type AssignOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignOperatorContext() *AssignOperatorContext {
	var p = new(AssignOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_assignOperator
	return p
}

func (*AssignOperatorContext) IsAssignOperatorContext() {}

func NewAssignOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignOperatorContext {
	var p = new(AssignOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_assignOperator

	return p
}

func (s *AssignOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(gengineParserASSIGN, 0)
}

func (s *AssignOperatorContext) SET() antlr.TerminalNode {
	return s.GetToken(gengineParserSET, 0)
}

func (s *AssignOperatorContext) PLUSEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserPLUSEQUAL, 0)
}

func (s *AssignOperatorContext) MINUSEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserMINUSEQUAL, 0)
}

func (s *AssignOperatorContext) MULTIEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserMULTIEQUAL, 0)
}

func (s *AssignOperatorContext) DIVEQUAL() antlr.TerminalNode {
	return s.GetToken(gengineParserDIVEQUAL, 0)
}

func (s *AssignOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAssignOperator(s)
	}
}

func (s *AssignOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAssignOperator(s)
	}
}

func (s *AssignOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAssignOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AssignOperator() (localctx IAssignOperatorContext) {
	localctx = NewAssignOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, gengineParserRULE_assignOperator)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(372)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-43)&-(0x1f+1)) == 0 && ((1<<uint((_la-43)))&((1<<(gengineParserASSIGN-43))|(1<<(gengineParserSET-43))|(1<<(gengineParserPLUSEQUAL-43))|(1<<(gengineParserMINUSEQUAL-43))|(1<<(gengineParserMULTIEQUAL-43))|(1<<(gengineParserDIVEQUAL-43)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IRangeOperatorContext is an interface to support dynamic dispatch.
type IRangeOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangeOperatorContext differentiates from other interfaces.
	IsRangeOperatorContext()
}

type RangeOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangeOperatorContext() *RangeOperatorContext {
	var p = new(RangeOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_rangeOperator
	return p
}

func (*RangeOperatorContext) IsRangeOperatorContext() {}

func NewRangeOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangeOperatorContext {
	var p = new(RangeOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_rangeOperator

	return p
}

func (s *RangeOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *RangeOperatorContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(gengineParserASSIGN, 0)
}

func (s *RangeOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangeOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangeOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterRangeOperator(s)
	}
}

func (s *RangeOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitRangeOperator(s)
	}
}

func (s *RangeOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitRangeOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) RangeOperator() (localctx IRangeOperatorContext) {
	localctx = NewRangeOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, gengineParserRULE_rangeOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(374)
		p.Match(gengineParserASSIGN)
	}

	return localctx
}

// INotOperatorContext is an interface to support dynamic dispatch.
type INotOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNotOperatorContext differentiates from other interfaces.
	IsNotOperatorContext()
}

type NotOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotOperatorContext() *NotOperatorContext {
	var p = new(NotOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_notOperator
	return p
}

func (*NotOperatorContext) IsNotOperatorContext() {}

func NewNotOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotOperatorContext {
	var p = new(NotOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_notOperator

	return p
}

func (s *NotOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *NotOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(gengineParserNOT, 0)
}

func (s *NotOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterNotOperator(s)
	}
}

func (s *NotOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitNotOperator(s)
	}
}

func (s *NotOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitNotOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) NotOperator() (localctx INotOperatorContext) {
	localctx = NewNotOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, gengineParserRULE_notOperator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(gengineParserNOT)
	}

	return localctx
}

// IMapVarContext is an interface to support dynamic dispatch.
type IMapVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMapVarContext differentiates from other interfaces.
	IsMapVarContext()
}

type MapVarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapVarContext() *MapVarContext {
	var p = new(MapVarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_mapVar
	return p
}

func (*MapVarContext) IsMapVarContext() {}

func NewMapVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapVarContext {
	var p = new(MapVarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_mapVar

	return p
}

func (s *MapVarContext) GetParser() antlr.Parser { return s.parser }

func (s *MapVarContext) AllVariable() []IVariableContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableContext)(nil)).Elem())
	var tst = make([]IVariableContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableContext)
		}
	}

	return tst
}

func (s *MapVarContext) Variable(i int) IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *MapVarContext) LSQARE() antlr.TerminalNode {
	return s.GetToken(gengineParserLSQARE, 0)
}

func (s *MapVarContext) RSQARE() antlr.TerminalNode {
	return s.GetToken(gengineParserRSQARE, 0)
}

func (s *MapVarContext) Integer() IIntegerContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerContext)
}

func (s *MapVarContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *MapVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapVarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterMapVar(s)
	}
}

func (s *MapVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitMapVar(s)
	}
}

func (s *MapVarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitMapVar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) MapVar() (localctx IMapVarContext) {
	localctx = NewMapVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, gengineParserRULE_mapVar)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(378)
		p.Variable()
	}
	{
		p.SetState(379)
		p.Match(gengineParserLSQARE)
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case gengineParserINT, gengineParserMINUS:
		{
			p.SetState(380)
			p.Integer()
		}

	case gengineParserDQUOTA_STRING:
		{
			p.SetState(381)
			p.StringLiteral()
		}

	case gengineParserSIMPLENAME, gengineParserDOTTEDNAME, gengineParserDOUBLEDOTTEDNAME:
		{
			p.SetState(382)
			p.Variable()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(385)
		p.Match(gengineParserRSQARE)
	}

	return localctx
}

// IAtNameContext is an interface to support dynamic dispatch.
type IAtNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtNameContext differentiates from other interfaces.
	IsAtNameContext()
}

type AtNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtNameContext() *AtNameContext {
	var p = new(AtNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atName
	return p
}

func (*AtNameContext) IsAtNameContext() {}

func NewAtNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtNameContext {
	var p = new(AtNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atName

	return p
}

func (s *AtNameContext) GetParser() antlr.Parser { return s.parser }
func (s *AtNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtName(s)
	}
}

func (s *AtNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtName(s)
	}
}

func (s *AtNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtName() (localctx IAtNameContext) {
	localctx = NewAtNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, gengineParserRULE_atName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)
		p.Match(gengineParserT__1)
	}

	return localctx
}

// IAtIdContext is an interface to support dynamic dispatch.
type IAtIdContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtIdContext differentiates from other interfaces.
	IsAtIdContext()
}

type AtIdContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtIdContext() *AtIdContext {
	var p = new(AtIdContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atId
	return p
}

func (*AtIdContext) IsAtIdContext() {}

func NewAtIdContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtIdContext {
	var p = new(AtIdContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atId

	return p
}

func (s *AtIdContext) GetParser() antlr.Parser { return s.parser }
func (s *AtIdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtIdContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtIdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtId(s)
	}
}

func (s *AtIdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtId(s)
	}
}

func (s *AtIdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtId() (localctx IAtIdContext) {
	localctx = NewAtIdContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, gengineParserRULE_atId)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(gengineParserT__2)
	}

	return localctx
}

// IAtDescContext is an interface to support dynamic dispatch.
type IAtDescContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtDescContext differentiates from other interfaces.
	IsAtDescContext()
}

type AtDescContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtDescContext() *AtDescContext {
	var p = new(AtDescContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atDesc
	return p
}

func (*AtDescContext) IsAtDescContext() {}

func NewAtDescContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtDescContext {
	var p = new(AtDescContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atDesc

	return p
}

func (s *AtDescContext) GetParser() antlr.Parser { return s.parser }
func (s *AtDescContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtDescContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtDescContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtDesc(s)
	}
}

func (s *AtDescContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtDesc(s)
	}
}

func (s *AtDescContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtDesc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtDesc() (localctx IAtDescContext) {
	localctx = NewAtDescContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, gengineParserRULE_atDesc)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(gengineParserT__3)
	}

	return localctx
}

// IAtSalContext is an interface to support dynamic dispatch.
type IAtSalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtSalContext differentiates from other interfaces.
	IsAtSalContext()
}

type AtSalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtSalContext() *AtSalContext {
	var p = new(AtSalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = gengineParserRULE_atSal
	return p
}

func (*AtSalContext) IsAtSalContext() {}

func NewAtSalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtSalContext {
	var p = new(AtSalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = gengineParserRULE_atSal

	return p
}

func (s *AtSalContext) GetParser() antlr.Parser { return s.parser }
func (s *AtSalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtSalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtSalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.EnterAtSal(s)
	}
}

func (s *AtSalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(gengineListener); ok {
		listenerT.ExitAtSal(s)
	}
}

func (s *AtSalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case gengineVisitor:
		return t.VisitAtSal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *gengineParser) AtSal() (localctx IAtSalContext) {
	localctx = NewAtSalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, gengineParserRULE_atSal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.Match(gengineParserT__4)
	}

	return localctx
}

func (p *gengineParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 9:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 10:
		var t *MathExpressionContext = nil
		if localctx != nil {
			t = localctx.(*MathExpressionContext)
		}
		return p.MathExpression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *gengineParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *gengineParser) MathExpression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
