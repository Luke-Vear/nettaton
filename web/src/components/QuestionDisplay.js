import React from 'react'

export const QuestionDisplay = ({ question, error }) =>
  <div className='question-display'>
    {contextMessage(question, error)}
  </div>

const contextMessage = (question, error) => {
  if (error) {
    return 'There appears to be something wrong with the server.'
  }
  if (question) {
    return questionKindMessage(question)
  }
  return 'Click next to get a new question.'
}

const questionKindMessage = (question) => {
  let message = 'What is the '

  switch (question.kind) {
    case 'first':
      message += 'first valid IP address'
      break
    case 'last':
      message += 'last valid IP address'
      break
    case 'broadcast':
      message += 'broadcast address'
      break
    case 'firstandlast':
      message += 'first and last valid IP addresses'
      break
    case 'hostsinnet':
      message += 'number of valid hosts'
      break
    default:
  }

  message += ' in the subnet that contains the address '
  message += question.ip + '/' + question.network

  return message + '?'
}
