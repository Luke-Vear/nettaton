import { connect } from 'react-redux'
import Display from '../components/Display'

const mapStateToProps = (state) => ({
  question: Display(state.question)
})

// const mapDispatchToProps = ({
//   onTodoClick: toggleTodo
// })

const DisplayQuestion = connect(
  mapStateToProps
//   mapDispatchToProps
)(Display)

export default DisplayQuestion