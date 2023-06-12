

describe('Rancher Web Page', () => {
  beforeEach(() => {
    cy.visit('https://localhost:8080/') 
  })

  it('should login into Rancher web page', () => {
    // Add test steps to login into Rancher web page
    cy.get('#username').type('admin') 
    cy.get('#password').type('newwork@1230@1A') 
    cy.get('button[type="submit"]').click() 

    // Assert that the login was successful
    cy.url().should('include', '/dashboard') 
  })

  it('should open the main web page', () => {
    // Assert that the main web page is opened
    cy.url().should('include', '/dashboard') 
  })

  it('should have the correct main web page title', () => {
    // Assert that the main web page title is correct
    cy.title().should('eq', 'Welcome to Rancher') 
  })
})
