import React from 'react'
import PropTypes from 'prop-types'

const Button = ({ onClick, text }) => (
  <li
    onClick={onClick}
    style={{
      color: 'red'
    }}
  >
    {text}
  </li>
)

Button.propTypes = {
  onClick: PropTypes.func.isRequired,
  text: PropTypes.string.isRequired
}

export default Button