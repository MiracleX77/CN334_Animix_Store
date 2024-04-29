import { configureStore, combineReducers, ThunkAction, Action } from '@reduxjs/toolkit';
import { persistReducer } from 'redux-persist';
import storage from 'redux-persist/lib/storage';
import localStorage from 'redux-persist/es/storage';

// modules
import exampleReducer from './modules/example/exampleSlice';


// config
const persistConfig = {
    key: 'root',
    storage: storage,
};

export const rootReducer: any = combineReducers({
    example: exampleReducer,
});

const persistedReducer: any = persistReducer(persistConfig, rootReducer);

export const makeStore = () => {
    return configureStore({
        reducer: persistedReducer,
        middleware: (getDefaultMiddleware) => getDefaultMiddleware(
            {
                serializableCheck: {
                    ignoredActions: ['persist/PERSIST'],
                },
            }
        ),
    });
};



export type RootState = ReturnType<typeof rootReducer>;
export type AppDispatch = ReturnType<typeof makeStore>['dispatch'];
export type AppStore = ReturnType<typeof makeStore>;
export type AppThunk = ThunkAction<void, RootState, unknown, Action<string>>;

export default makeStore;
