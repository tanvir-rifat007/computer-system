#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <string.h>

// Check if byte is start of 2-byte sequence
bool is_len_two(uint8_t b) {
    return b >= 0xC0 && b <= 0xDF;
}

// Check if byte is start of 3-byte sequence
bool is_len_three(uint8_t b) {
    return b >= 0xE0 && b <= 0xEF;
}

// Check if byte is start of 4-byte sequence
bool is_len_four(uint8_t b) {
    return b >= 0xF0 && b <= 0xF7;
}

int main() {
    FILE *fin = fopen("cases", "rb");
    FILE *fout = fopen("out", "w");

    if (!fin || !fout) {
        perror("Failed to open file");
        return 1;
    }

    uint8_t buffer[1024];
    size_t line_num = 0;

    while (fgets((char *)buffer, sizeof(buffer), fin)) {
        line_num++;

        size_t line_len = strlen((char *)buffer);
        if (line_len < 2) continue;

        uint8_t trunc_len = buffer[0];        // first byte is truncate length
        uint8_t *data = &buffer[1];           // actual utf-8 data starts here
        size_t data_len = line_len - 2;       // minus trunc_len + newline

        size_t out_len = 0;
        for (size_t i = 0; i < data_len;) {
            if (out_len >= trunc_len) break;

            uint8_t byte = data[i];
            size_t char_len = 1;

            if (is_len_four(byte))      char_len = 4;
            else if (is_len_three(byte)) char_len = 3;
            else if (is_len_two(byte))   char_len = 2;

            if (out_len + char_len > trunc_len) break;

            fwrite(&data[i], 1, char_len, fout);
            i += char_len;
            out_len += char_len;
        }

        fputc('\n', fout);
    }

    fclose(fin);
    fclose(fout);
    return 0;
}

