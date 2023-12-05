#include <stdio.h>
#include <string.h>

#define LPAREN   '('
#define RPAREN   ')'
#define LBRACE   '['
#define RBRACE   ']'
#define LBRACKET '{'
#define RBRACKET '}'
#define LCARET   '<'
#define RCARET   '>'

static const char pairs[4][2] = {
    { LPAREN, RPAREN },
    { LBRACE, RBRACE },
    { LBRACKET, RBRACKET },
    { LCARET, RCARET },
};

static const size_t scores[4] = {
    3,
    57,
    1197,
    25137
};

static const size_t costs[4] = {
    1,
    2,
    3,
    4
};

static inline char partner(const char ch) {
    for (int index = 0; index < sizeof(pairs); index++) {
        if (ch == pairs[index][0]) { return pairs[index][1]; }
        if (ch == pairs[index][1]) { return pairs[index][0]; }
    }
    return -1;
}

static inline int index_of(const char ch)
{
    switch (ch) {
    case LPAREN:
    case RPAREN:
        return 0;
    case LBRACE:
    case RBRACE:
        return 1;
    case LBRACKET:
    case RBRACKET:
        return 2;
    case LCARET:
    case RCARET:
        return 3;
    default:
        return -1;
    }
}

static inline int direction(const char ch)
{
    switch (ch) {
    case LPAREN:
    case LBRACE:
    case LBRACKET:
    case LCARET:
        return 1;
    case RPAREN:
    case RBRACE:
    case RBRACKET:
    case RCARET:
        return -1;
    default:
        return 0;
    }
}

static size_t find_error(const char* line)
{
    char buf[1024] = {0};
    char *ptr = buf;
    size_t num_chars = strlen(line);

    for (int i = 0; i < num_chars; i++) {
        char ch = line[i];
        int dir = direction(ch);
        if (dir > 0) {
            *ptr++ = ch;
        } else if (dir < 0 && ptr > buf) {
            ptr--;
            char expected = partner(*ptr);
            if (expected != ch) {
                //printf("%s - Expected %c, but found %c instead.\n", line, expected, ch);
                return scores[index_of(ch)];
            }
        }
    }
    return 0;
}

static size_t fix_incomplete(const char* line)
{
    char buf[1024] = {0};
    char *ptr = buf;
    size_t num_chars = strlen(line);

    for (int i = 0; i < num_chars; i++) {
        char ch = line[i];
        int dir = direction(ch);
        if (dir > 0) {
            *ptr++ = ch;
        } else if (dir < 0 && ptr > buf) {
            ptr--;
            char expected = partner(*ptr);
            if (expected != ch) {
                //printf("%s - Expected %c, but found %c instead.\n", line, expected, ch);
                return 0;
            }
        }
    }

    size_t cost = 0;
    while (--ptr >= buf) {
        cost *= 5;
        cost += costs[index_of(*ptr)];
    }
    return cost;
}

int main(int argc, char **argv)
{
    size_t cost = 0;
    char line[256 + 1];
    while (fgets(line, 256, stdin)) {
        line[strlen(line) - 1] = 0;
        cost = fix_incomplete(line);
        if (cost > 0) {
            printf("%zu\n", cost);
        }
    }
    return 0;
}
