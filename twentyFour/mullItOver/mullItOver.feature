Feature: Mull it over

  Scenario Outline: Shopkeepers computer
    Given There is corrupted <input>
    When We parse the input there should be <amount_of_mods>
    When we add up these results
    Then they should be <result>
    Examples:
      | title       | input                             | amount_of_mods | result |
      | basic       | "ct():@mul(56,920)-how()}"        | 1              | 51520  |
      | four digits | "ct():@mul(56,9200)-how()}"       | 1              | 515200 |
      | two basics  | "ct():@mul(56,9)-homul(56,9)w()}" | 2              | 1008   |
