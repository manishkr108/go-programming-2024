Golang Array vs Slice

Feature                                     Array                                                       Slice
Declaration                    var arr[length]type                                          var slice []type
                            myArray := [...]int{1,2,3}                                      slice :=make([]type, length)

Length                      fixed at declaration                                            Dynamic, can grow or shrink

underlying                  Contiguous block of memory                                      Referance to an underlying array


