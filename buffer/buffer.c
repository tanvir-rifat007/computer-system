// first create a file like : touch /tmp/foo
// and then run : tail -f /tmp/foo 
// if we don't flush the file, we won't see the data written to it live.


#include<stdio.h>
#include<stdlib.h>
#include<string.h>



int main(){

    FILE *fp = fopen("/tmp/foo", "w");

    if (!fp) {
        perror("Failed to open file");
        return 1;
    }

    const char *data = "Hello, world! This is a test string for truncation.";
    size_t data_len = strlen(data);
    
    // Write the data to the file

    fwrite(data, sizeof(char), data_len, fp); 

    fclose(fp);
    fflush(fp);


 return 0;
}
