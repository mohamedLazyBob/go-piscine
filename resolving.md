#   ALL STEP FOR RESOLVING THE CRIME
##  STEP : 1
> Collect all the clues at the crime scene

```sh
grep "CLUE" crimescene
```

### **RESULT**

**CLUE**: Footage from an ATM security camera is blurry but shows that the perpetrator is a **tall male**, **at least 6**'.
**CLUE**: Found a wallet believed to belong to the killer: no ID, just loose change, and **membership cards for AAA**, **Delta SkyMiles, the local library, and the Museum of Bash History**. The cards are totally untraceable and have no name, for some reason.
**CLUE**: Questioned the barista at the local coffee shop. He said a woman left right before they heard the shots. The name on her latte was **Annabel**, she had **blond spiky hair** and a **New Zealand** accent.

### **INTERPRETATION**
+   The murder is a tall male
+   The murder is at leat 6 tall
+   The murder have a membership cards for AAA
+   The barista saw Annabel (a blond spiky hair at New Zealand)


-

##  STEP : 2
> Investigating about who is Annabel

###   Get all people who is named Annabel

```sh
grep "Annabel" people
```

### **RESULT**

**Annabel** Sun     F       26      Hart Place, line 40
Oluwasegun **Annabel**      M       37      Mattapan Street, line 173
**Annabel** Church  F       38      Buckingham Place, line 179
**Annabel** Fuglsang        M       40      Haley Street, line 176

### **INTERPRETATION**
+   Annabel Sun is a female
+   Annabel Church is a female

###   Get the correct interviews key for Annabel Sun

```sh
head -n 40 streets/Hart_Place | tail -n 1
```

### **RESULT**
SEE INTERVIEW #47246024


###   Get interview 47246024

```sh
cat interviews/interview-47246024
```

### **RESULT**

**Ms. Sun** has **brown hair** and is not from New Zealand.  Not the witness from the cafe.

### **INTERPRETATION**
+   Annabel Sun is brown hair
+   We must check the interview of Annabel Church

###   Get the correct interviews key for Annabel Church

```sh
head -n 179 streets/Buckingham_Place | tail -n 1
```

### **RESULT**
SEE INTERVIEW #699607


###   Get interview 47246024

```sh
cat interviews/interview-699607
```

### **RESULT**

Interviewed Ms. **Church** at 2:04 pm.  Witness stated that she did **not see anyone** she could identify as the shooter, that she ran away as soon as the shots were fired.

However, she reports seeing the car that fled the scene.  Describes it as a **blue Honda**, with a license plate that starts with **"L337"** and ends with **"9"**

### **INTERPRETATION**
+   Annabel Sun did not see anyone
+   Annabel see a blue honda with license plate starts with "L337" and ends with "9"


-

##  STEP 3 :
> Investing about the car annabel saw

### Get the list of all car that is a blue honda with a license plate starting with L337

```sh
grep -A 5 "L337" vehicles | grep -B 1 -A 5 "Honda" | grep -A 5 -B 2 "Blue"
```

### **RESULT**

License Plate **L337**QE9
Make: **Honda**
Color: **Blue**
Owner: Erika Owens
Height: 6'5"
Weight: 220 lbs

License Plate **L337**539
Make: **Honda**
Color: **Blue**
Owner: Aron Pilhofer
Height: 5'3"
Weight: 198 lbs

License Plate **L337**369
Make: **Honda**
Color: **Blue**
Owner: Heather Billings
Height: 5'2"
Weight: 244 lbs

License Plate **L337**DV9
Make: **Honda**
Color: **Blue**
Owner: Joe Germuska
Height: 6'2"
Weight: 164 lbs

License Plate **L337**5A9
Make: **Honda**
Color: **Blue**
Owner: Dartey Henv
Height: 6'1"
Weight: 204 lbs

License Plate **L337**WR9
Make: **Honda**
Color: **Blue**
Owner: Hellen Maher
Height: 6'2"
Weight: 130 lbs

### **INTERPRETATION**
+   We have user with a height lower than the suspect

### refine the search by filtering by height

```sh
grep -A 5 "L337" vehicles | grep -B 1 -A 5 "Honda" | grep -A 5 -B 2 "Blue"
```

### **RESULT**

License Plate **L337**QE9
Make: **Honda**
Color: **Blue**
Owner: Erika Owens
**Height: 6**'5"
Weight: 220 lbs

License Plate **L337**DV9
Make: **Honda**
Color: **Blue**
Owner: Joe Germuska
**Height: 6**'2"
Weight: 164 lbs

License Plate **L337**5A9
Make: **Honda**
Color: **Blue**
Owner: Dartey Henv
**Height: 6**'1"
Weight: 204 lbs

License Plate **L337**WR9
Make: **Honda**
Color: **Blue**
Owner: Hellen Maher
**Height: 6**'2"
Weight: 130 lbs

### Look for the suspects membership list

```sh
cat memberships/Delta_SkyMiles memberships/AAA memberships/Museum_of_Bash_History memberships/Terminal_City_Library | grep -c "Erika Owens"
```

### **RESULT**
0

```sh
cat memberships/Delta_SkyMiles memberships/AAA memberships/Museum_of_Bash_History memberships/Terminal_City_Library | grep -c "Joe Germuska"
```

### **RESULT**
2

```sh
cat memberships/Delta_SkyMiles memberships/AAA memberships/Museum_of_Bash_History memberships/Terminal_City_Library | grep -c "Dartey Henv"
```

### **RESULT**
4

```sh
cat memberships/Delta_SkyMiles memberships/AAA memberships/Museum_of_Bash_History memberships/Terminal_City_Library | grep -c "Hellen Maher"
```

### **RESULT**
4

### **INTERPRETATION**
+   Dartey Henv have all memberships
+   Hellen Maher have all memberships
+   Dartey Henv is a male because of his name

#   🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩
#   🚩🚩🚩DARTEY HENV IS THE MURDER🚩🚩🚩
#   🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩🚩