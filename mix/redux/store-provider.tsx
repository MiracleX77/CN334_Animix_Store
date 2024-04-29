'use client'
import { useRef } from 'react';
import { Provider } from 'react-redux';
import { makeStore } from './index';
import { PersistGate } from 'redux-persist/integration/react';
import { persistStore, Persistor } from 'redux-persist'; // Import Persistor type from redux-persist

import { Store } from 'redux'; // Import Store type from redux

export default function StoreProvider({ children }: Readonly<{ children: React.ReactNode }>) {
    const storeRef = useRef<Store>(); // Specify the type for storeRef
    const persistor = useRef<Persistor | undefined>(undefined); // Specify the type for persistor

    if (!storeRef.current) {
        // Create the store instance the first time this renders
        storeRef.current = makeStore();
        persistor.current = persistStore(storeRef.current);
    }
    return (
        <Provider store={storeRef.current}>
            <PersistGate loading={null} persistor={persistor.current as Persistor}>
                {children}
            </PersistGate>
        </Provider>
    );
}
