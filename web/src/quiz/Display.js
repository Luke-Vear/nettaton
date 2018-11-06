import React from 'react'

export const Display = ({ answer, busy, result, question, error }) => {
  let message = (() => {
    if (answer) {
      return answer
    }
    if (busy) {
      return 'Fetching...'
    } else if (result) {
      return result ? 'Correct!' : 'Wrong!'
    } else if (question) {
      return questionKindMessage(question)
    } else if (error) {
      return 'There appears to be something wrong with the server.'
    } else {
      return 'Click below to begin'
    }
  })()

  return (
    <div>
      <p>{message}</p>
    </div>
  )
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
