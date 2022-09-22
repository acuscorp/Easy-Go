package cmp_test

import (
  "testing"
   mycmp "test_examples/cmp"
  "github.com/google/go-cmp/cmp"
)

func TestCreatePerson(t *testing.T) {
  expected := mycmp .Person {
    Name: "Dennis",
    Age: 37,
  }
  result := mycmp .CreatePerson("Dennis", 37)

  if diff := cmp.Diff(expected, result); diff != "" {
    t.Error(diff)
  } 
}

 func TestCreatePersonIgnoreDate(t *testing.T) {
   expected := mycmp.Person{
     Name: "Dennis",
     Age: 37,
   }

   result := mycmp.CreatePerson("Dennis", 37)
   comparer := cmp.Comparer(func(x,y mycmp.Person) bool {
     return x.Name == y.Name && x.Age == y.Age
   })

   if diff := cmp.Diff(expected, result,comparer); diff != "" {
     t.Error(diff)
   }

   if result.DateAdded.IsZero() {
     t.Error("DateAdded wasn't assigned")
   }
 }