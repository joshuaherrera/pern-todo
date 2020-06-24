/* eslint-disable react/jsx-filename-extension */
import React, { Fragment } from 'react';
import './App.css';

// components
import InputTodo from './components/InputTodo';

function App() {
  return (
    // eslint-disable-next-line react/jsx-fragments
    <Fragment>
      <div className="container">
        <InputTodo />

      </div>
    </Fragment>
  );
}

export default App;
