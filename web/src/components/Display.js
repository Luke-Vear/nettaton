import React from 'react'
import PropTypes from 'prop-types'

const Display = ({ question }) => (
  <p>
    {question}
  </p>
)

Display.propTypes = {
  question: PropTypes.shape({
    id: PropTypes.string.isRequired,
    ip: PropTypes.string.isRequired,
    network: PropTypes.string.isRequired,
    kind: PropTypes.string.isRequired,
    ttl: PropTypes.number.isRequired
  }).isRequired
}

export default Display

// {
//     "id": "b1457fd8-972f-4de6-8ca6-8653bd65139d",
//     "ip": "172.17.71.145",
//     "network": "19",
//     "kind": "first",
//     "ttl": 1540493304
// }