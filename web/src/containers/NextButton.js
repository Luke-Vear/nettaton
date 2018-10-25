import { connect } from 'react-redux'
import { fetchQuestion } from '../actions'
import Button from '../components/Button'

const mapStateToProps = (state) => ({
  text: 'sometext'
})

const mapDispatchToProps = (dispatch) => ({
  onClick: () => {
    dispatch(fetchQuestion())
  }
})

const NextButton = connect(
  mapStateToProps,
  mapDispatchToProps
)(Button)

export default NextButton
