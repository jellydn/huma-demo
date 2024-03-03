import { expect } from 'chai';
import { describe, it } from 'mocha';
import { myBusinessLogic } from './index';

describe('myBusinessLogic', () => {
  it('should log "Hello via Bun!"', () => {
    // Mock console.log
    const consoleLogMock = jest.spyOn(console, 'log');
    
    // Call the business logic function
    myBusinessLogic();
    
    // Verify that console.log was called with the expected message
    expect(consoleLogMock).to.have.been.calledWith('Hello via Bun!');
    
    // Restore the original console.log function
    consoleLogMock.mockRestore();
  });
});
