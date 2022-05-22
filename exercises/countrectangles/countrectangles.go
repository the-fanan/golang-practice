package rectangles

/**
Count the rectangles in an ASCII diagram like the one below.


   +--+
  ++  |
+-++--+
|  |  |
+--+--+

Approach:

1. We notice that the array of strings forms a matrix
2. We loop through the each cell in the matrix till we find a CORNER
3. Once a corner is found, count number of  rectangles corner forms
4. Add this to total rectangles count

For counting number of rectangles a corner forms,
5. we check from it's right neighbour
6. We then follow the rules edges and corners of the rectangle in the order RIGHT, DOWN, LEFT, UP
7. if we encounter an edge and it matches the direction being traversed, we keep exploring to find if it forms a valid rectangle
8. If we encounter a corner and incoming direction is UP. if corner == origin the return 1 else keep following incoming direction to check if valid rectangle is found
9. If we encounter a corner and incoming direction is not UP we check for valid rectangle by following the current direction and the next direction in the specified direction order in step 6
10. If the ro or column is out of range return 0


Time Complexity O((N*M)^2) worst case when every cell is a CORNER
Space Complexity O(N*M) worst case depth of stack when every cell is a CORNER
**/
const RIGHT string = "right"
const DOWN string = "down"
const LEFT string = "left"
const UP string = "up"
const CORNER byte = '+'
const HORIZONTAL_EDGE byte = '-'
const VERTICAL_EDGE byte = '|'

func Count(diagram []string) int {
    rectangles := 0
	rows := len(diagram)
    if rows == 0 {
        return rectangles
    }
	cols := len(diagram[0])
    for row := 0; row < rows; row++ {
        for col := 0; col < cols; col++ {
            if diagram[row][col] == CORNER {
				count := countRectangles(diagram, RIGHT, []int{row, col}, row, col + 1)
                rectangles += count
            }
        }
    }

    return rectangles
}

func countRectangles(diagram []string, direction string, origin []int, row, col int) int {
	rectangles := 0
	rows := len(diagram)
	cols := len(diagram[0])
	if 0 > row || row >= rows || 0 > col || col >= cols {
		return 0
	}
	// direction order: right, down, left, up
	if diagram[row][col] == CORNER {
		if direction == UP {
			if row == origin[0] && col == origin[1] {
				return 1
			} else {
				rectangles += countRectangles(diagram, UP, origin, row - 1, col)
			}
		}

		if direction == RIGHT {
			rectangles += countRectangles(diagram, DOWN, origin, row + 1, col)
			rectangles += countRectangles(diagram, RIGHT, origin, row, col + 1)
		}

		if direction == DOWN {
			rectangles += countRectangles(diagram, LEFT, origin, row, col - 1)
			rectangles += countRectangles(diagram, DOWN, origin, row + 1, col)
		}

		if direction == LEFT {
			rectangles += countRectangles(diagram, UP, origin, row - 1, col)
			rectangles += countRectangles(diagram, LEFT, origin, row, col - 1)
		}
	} else if diagram[row][col] == HORIZONTAL_EDGE {
		if direction == LEFT{
			rectangles += countRectangles(diagram, LEFT, origin, row, col - 1)
		}
		if direction == RIGHT {
			rectangles += countRectangles(diagram, RIGHT, origin, row, col + 1)
		}
	} else if diagram[row][col] == VERTICAL_EDGE {
		if direction == UP{
			rectangles += countRectangles(diagram, UP, origin, row - 1, col)
		}
		if direction == DOWN {
			rectangles += countRectangles(diagram, DOWN, origin, row + 1, col)
		}
	}
	return rectangles
}