#! /bin/bash

# Run in folder below /mystery

# Define variables
witness="Annabel Church"
car_make="Honda"
car_color="Blue"
license="L337*"
height="6"
main_suspect="Bad Dude"

# Search for the witness in people
output=$(grep "$witness" mystery/people)

# Error and exit if nothing found
if [ -z "$output" ]; then
    echo "No matches found for '$witness' in mystery/people."
    exit 1
fi

# Get the street name (fields from 5th to 3rd to last) using awk and sed
street=$(echo "$output" | awk '{for (i=5; i<=(NF-2); i++) printf "%s ", $i}' | sed 's/,//g' | sed 's/ /_/g') 
# Remove the last character
street="${street%?}"

# Get the address at the end of the line using awk
address=$(echo "$output" | awk '{print $NF}')

# Interview number from street and address
interview_number=$(cat mystery/streets/"$street" | head -n +"$address" | tail -n 1 | cut -d "#" -f2)

# Make the environement variable
#export IV_NUM=$interview_number
export IV_NUM=$699607
echo ${IV_NUM}

# Print out the interview
#cat mystery/interviews/interview-$interview_number
cat mystery/interviews/interview-$699607
echo ""

grep -A 5 "$license" mystery/vehicles | grep -B 1 -A4 "$car_make"| grep -B 2 -A 3 "$car_color"  | grep -A 1 -B 4 "Height: $height"

# Print out the contents of the environement variable as requested
export MAIN_SUSPECT=$main_suspect
echo ""
echo "Main suspect:"
echo ${MAIN_SUSPECT}