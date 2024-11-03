#include <stdio.h>
#include <inttypes.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

#define MAX_LINE_LENGTH 256

typedef struct point_t {
    unsigned x;
    unsigned y;
} point_t;

typedef struct line_t {
    point_t p1;
    point_t p2;
} line_t;

typedef struct list_t {
    void* data;
    size_t unit;
    size_t capacity;
    size_t length;
} list_t;

static inline size_t list_size(const list_t* list)
{
    return list->unit * list->length;
}

static inline list_t new_list(size_t unit, size_t capacity)
{
    list_t list = { .data = NULL, .unit = unit, .capacity = capacity, .length = 0 };
    list.data = (void*)malloc(list.unit * list.capacity);
    if (!list.data) {
        perror("Error: ");
        exit(1);
    }
    return list;
}

static inline void list_grow(list_t* list, size_t factor)
{
    void *swap = list->data;
    void *copy = malloc(list_size(list) * factor);
    if (!copy) {
        perror("Error: ");
        exit(1);
    }
    memcpy(copy, list->data, list_size(list));
    list->data = copy;
    list->capacity *= factor;
    free(swap);
}

typedef struct grid_t {
    void* data;
    size_t unit;
    size_t width;
    size_t height;
} grid_t;

static inline void* grid_at(const grid_t* grid, size_t x, size_t y)
{
    return &grid->data[grid->unit * ((y * grid->width) + x)];
}

static inline size_t grid_size(const grid_t* grid)
{
    return grid->unit * grid->height * grid->width;
}

static inline grid_t new_grid(size_t unit, size_t width, size_t height)
{
    grid_t grid = { .data = NULL, .unit = unit, .width = width, .height = height };
    grid.data = (void*)malloc(grid_size(&grid));
    if (!grid.data) {
        perror("Error: ");
        exit(1);
    }
    return grid;
}

static list_t parse_lines(FILE* file, size_t* height, size_t* width)
{
    char input[256 + 1] = {0};
    size_t w = 0, h = 0;
    list_t lines = new_list(sizeof(line_t), 10);
    while (fgets(input, MAX_LINE_LENGTH, file)) {
        line_t line;
        if (sscanf(input, "%u,%u -> %u,%u\n", &line.p1.x, &line.p1.y, &line.p2.x, &line.p2.y) < 4) {
            fprintf(stderr, "Error: parsed unexpected number of values from input on line %lu:\n'%s'", lines.length + 1, input);
            exit(1);
        }
        if (lines.length == lines.capacity) {
            list_grow(&lines, 2);
        }
        ((line_t*)lines.data)[lines.length++] = line;
        if (line.p1.x > w)  { w = line.p1.x; }
        if (line.p2.x > w)  { w = line.p2.x; }
        if (line.p1.y > h)  { h = line.p1.y; }
        if (line.p2.y > h)  { h = line.p2.y; }
    }
    *width = w + 1;
    *height = h + 1;
    return lines;
}

static grid_t fill_grid(list_t* lines, size_t height, size_t width)
{
    grid_t grid = new_grid(sizeof(unsigned), height, width);
    for (int i = 0; i < lines->length; i++) {
        line_t *line = &((line_t*)lines->data)[i];
        if (line->p1.x == line->p2.x) {
            // vertical
            const point_t *upper, *lower;
            if (line->p1.y < line->p2.y) { upper = &line->p1; lower = &line->p2; }
            else                         { upper = &line->p2; lower = &line->p1; }
            for (int y = upper->y; y <= lower->y; y++) {
                ++(*(unsigned*)grid_at(&grid, upper->x, y));
            }
        } else if (line->p1.y == line->p2.y) {
            // horizontal
            const point_t *left, *right;
            if (line->p1.x < line->p2.x) { left = &line->p1; right = &line->p2; }
            else                         { left = &line->p2; right = &line->p1; }
            for (int x = left->x; x <= right->x; x++) {
                ++(*(unsigned*)grid_at(&grid, x, left->y));
            }
        } else {
            // diagonal 
            const point_t *upper, *lower;
            if (line->p1.y < line->p2.y) { upper = &line->p1; lower = &line->p2; }
            else                         { upper = &line->p2; lower = &line->p1; }
            if (upper->x < lower->x) {
                // right
                int y, x;
                for (y = upper->y, x = upper->x; y <= lower->y; y++, x++) {
                    ++(*(unsigned*)grid_at(&grid, x, y));
                }
            } else {
                // left
                int y, x;
                for (y = upper->y, x = upper->x; y <= lower->y; y++, x--) {
                    ++(*(unsigned*)grid_at(&grid, x, y));
                }
            }
        }
    }
    return grid;
}

static void draw_grid(const grid_t* grid)
{
    printf("(%lu, %lu)\n", grid->width, grid->height);
    for (int y = 0; y < grid->height; y++) {
        for (int x = 0; x < grid->width; x++) {
            unsigned value = *(unsigned*)grid_at(grid, x, y);
            if (value > 0) {
                printf("%u", value);
            } else {
                printf(".");
            }
        }
        printf("\n");
    }
}

static size_t calculate_num_overlapping_points(const grid_t* grid)
{
    size_t count = 0;
    for (int y = 0; y < grid->height; y++) {
        for (int x = 0; x < grid->width; x++) {
            if (*(unsigned*)grid_at(grid, x, y) > 1) {
                count++;
            }
        }
    }
    return count;
}

int main(int argc, char **argv)
{
    size_t width = 0, height = 0;
    list_t lines = parse_lines(stdin, &height, &width);
    grid_t grid = fill_grid(&lines, height, width);
    //draw_grid(&grid);
    printf("%lu\n", calculate_num_overlapping_points(&grid));
}
