#include<stdio.h>
#include<stdlib.h>
#include<ctype.h>

#define ALPHABET_SIZE 26

int isPangram(char *str){
  int seen[ALPHABET_SIZE] = {0};

  for(int i =0; str[i] != '\0'; i++){
    if(isalpha(str[i])){
      char lowerChar = tolower(str[i]);
      seen[lowerChar - 'a'] = 1; // Mark the letter as seen

      // this convert like:

      // seen['a' - 'a'] = 0 ; for a
      // seen['b' - 'a'] = 1 ; for b
      // seen['c' - 'a'] = 2 ; for c
      // seen['z' - 'a'] = 25 ; for z
    }

  }

  for(int i = 0;i<ALPHABET_SIZE;i++){
    if(seen[i] == 0){
      return 0;
    }

  }

  return 1;
}


int main(){
  char input[1000];
  printf("Enter a string: ");
  fgets(input, sizeof(input),stdin);

  int res = isPangram(input);
  if(res){
    printf("The string is a pangram.\n");
  } else {
    printf("The string is not a pangram.\n");
  }




}