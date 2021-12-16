import React from 'react';
import { BrowserRouter } from 'react-router-dom';
import { AppRouter } from './router/Router';
import { theme } from './theme/theme';
import { ChakraProvider } from "@chakra-ui/react"
import { Provider } from "react-redux";
import createInitStore from './reducks/store/store';

// storeをインスタンス
const store = createInitStore()

function App() {
  return (
    <>
    <ChakraProvider theme={theme}>
      {/* reactとreduxを接続してstoreを変更できるようにする */}
      <Provider store={store}>
        <BrowserRouter>
          <AppRouter/>
        </BrowserRouter>
      </Provider>
    </ChakraProvider>
    </>
  );
}

export default App;