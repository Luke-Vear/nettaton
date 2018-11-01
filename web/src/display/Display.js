import React from 'react'
import { connect } from 'react-redux'

export const Display = ({ fetching, question, error }) => {
  let message = 'Click below to begin'

  if (fetching) {
    message = 'Fetching...'
  } else if (question) {
    message = questionKindMessage(question)
  } else if (error) {
    message = 'There appears to be something wrong with the server.'
  }

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

const mapStateToProps = state => ({
  question: state.question,
  fetching: state.fetching,
  error: state.error
})

export default connect(mapStateToProps)(Display)
