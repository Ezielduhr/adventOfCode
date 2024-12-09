Feature: historian hysteria
  In order to find the chief historian
  We need to have a list with locations he visited
  These two list needs to be reconciled

  Scenario Outline: total distance between lists
    Given there are <firstList> and <secondList> to compare
    When I compare the lists for total distance
    Then the totalDistance should equal <totalDistance>
    Examples:
      | title   | firstList | secondList | totalDistance |
      | 'left'  | 1 2 3 4 5 | 2 3 4 5 6  | 5             |
      | 'right' | 2 3 4 5 6 | 1 2 3 4 5  | 5             |
      | 'both'  | 1 3 4 5 6 | 2 4 5 6 7  | 5             |
      | 'sort'  | 4 5 3 2 1 | 2 5 3 6 4  | 5             |

  Scenario Outline: calculate similarity score
    Given there are <firstList> and <secondList> to compare
    When I compare the lists for similarity
    Then the similarity should equal <similarity>
    Examples:
      | title      | firstList  | secondList | similarity |
      | 'multiple' | 1 2 3 4 5  | 1 1 1 5 6  | 4          |
      | 'numbers'  | 20 3 4 5 6 | 1 20 3 4 5 | 4          |
