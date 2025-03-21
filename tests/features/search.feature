Feature: Search and view person details

  Scenario: User searches for a person and views their detail page
    Given the server is running
    When I search for "Alice"
    And I click on the first result
    Then I should see the person's name as "Alice"
