import React, { Component } from 'react'
import logo from './logo.svg'
import './App.css'
import { apiCallRequest } from './redux'

import { connect } from 'react-redux'

class App extends Component {
  render () {
    const { fetching, dog, onRequestDog, error } = this.props

    return (
      <div className='App'>
        <header className='App-header'>
          <img src={dog || logo} className='App-logo' alt='logo' />
          <h1 className='App-title'>Welcome to Dog Saga</h1>
        </header>

        {dog ? (
          <p className='App-intro'>Keep clicking for new dogs</p>
        ) : (
          <p className='App-intro'>Replace the React icon with a dog!</p>
        )}

        {fetching ? (
          <button disabled>Fetching...</button>
        ) : (
          <button onClick={onRequestDog}>Request a Dog</button>
        )}

        {error && <p style={{ color: 'red' }}>Uh oh - something went wrong!</p>}

        <div style={{ display: 'block', padding: '50px 0px 10px 0px' }}>
          <a
            href='https://hackernoon.com/redux-saga-tutorial-for-beginners-and-dog-lovers-aa69a17db645'
            target='_blank'
            rel='noopener noreferrer'
            style={{ display: 'block', padding: '20px 5%' }}
          >
            Read the tutorial here
          </a>
          <a
            style={{ padding: '0px 5%' }}
            className='github-button'
            href='https://github.com/ryanjyost/redux-saga-beginner-tutorial'
            data-icon='octicon-star'
            aria-label='Star ryanjyost/dom-events on GitHub'
          >
            Star
          </a>
        </div>

        <div style={{ display: 'block', padding: '10px 5%' }}>
          <a
            style={{ padding: '0px 5%' }}
            href='https://twitter.com/share?ref_src=twsrc%5Etfw'
            className='twitter-share-button'
            data-text='Check out Dog Saga - A Redux-Saga Beginner Tutorial'
            data-url='https://ryanjyost.github.io/redux-saga-beginner-tutorial/'
            data-related='ryanjyost'
            data-show-count='false'
          >
            Tweet
          </a>
        </div>
        <div style={{ display: 'block', padding: '10px 5%' }}>
          <a
            style={{ padding: '0px 5%' }}
            href='https://twitter.com/ryanjyost?ref_src=twsrc%5Etfw'
            className='twitter-follow-button'
            data-show-count='false'
          >
            Follow @ryanjyost
          </a>
        </div>
      </div>
    )
  }
}

const mapStateToProps = state => {
  return {
    fetching: state.fetching,
    dog: state.dog,
    error: state.error
  }
}

const mapDispatchToProps = dispatch => {
  return {
    onRequestDog: () => dispatch(apiCallRequest())
  }
}

export default connect(mapStateToProps, mapDispatchToProps)(App)
