Feature: Red-Nosed Reports
  In order to find the chief historian
  We need to analyse some unusual data (while they search the reactor)
  These reports need to be analysed for safe levels
#  TODO this should have included a column with wath checked failed

  Scenario Outline: safety levels
    Given there are safety <levels>
    When we check the report
    Then the report should be <result>
    Examples:
      | title          | levels      | result |
      | increasing     | 1 3 6 7 9   | safe   |
      | direction      | 3 1 4 7 9   | unsafe |
      | decreasing     | 7 6 4 2 1   | safe   |
      | increasing dip | 1 3 6 9 7   | unsafe |
      | four increase  | 1 5 8 9 10  | unsafe |
      | stale          | 1 3 3 6 7 8 | unsafe |

  Scenario Outline: problem dampener
    Given there are safety <levels>
    Given we enable the problem dampener
    When we check the report
    Then the report should be <result>
    Examples:
      | title                 | levels         | result |
      | direction             | 3 1 2 5 7 9    | safe   |
      | increasing            | 1 3 2 5 7 9    | safe   |
      | na direction          | 3 1 2 5 4 6    | unsafe |
      | non adjacent increase | 1 3 2 5 2 7 9  | unsafe |
      | decreasing            | 7 6 4 2 1      | safe   |
      | increasing dip        | 1 3 6 9 7 10   | safe   |
      | double rapid increase | 1 5 9 10 12 13 | unsafe |
      | na rapid increase     | 1 5 8 9 10 15  | unsafe |
      | staleSafe             | 1 3 3 6 7 8    | safe   |
      | non adjacent stale    | 1 3 6 3 7 8    | safe   |
      | staleUnsafe           | 1 3 3 3 6 7 8  | unsafe |
      | last                  | 1 3 6 8 8      | safe   |