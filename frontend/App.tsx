import { StatusBar } from 'expo-status-bar';
import React, {useState, useEffect, Component} from 'react';
import { Alert, Button,TouchableOpacity, StyleSheet, Text, View } from 'react-native';
import TaskCard from './components/TaskCard';
import { useTasks } from './hooks/useTasks'



export default function App() {
  const {data, loading, error} = useTasks();
  if (loading) {
    return(
      <View style={styles.container}>
        <Text style={{backgroundColor:"lightgreen"}}> WAITING BRUH... 🕓</Text>
      </View>
    );
  }
  if (error) {
    return(
      <View style={styles.container}>
        <Text style={{backgroundColor:"lightgreen"}}> WHAT'S THAT BRUH... 🤮</Text>
      </View>
    );
  }
  return (
    <View style={styles.container}>
      {data.map((task,index) => (
        <View key={index}>
          <TaskCard title={task.title} completed={task.completed}/>
        </View>
      ))}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#000',
    alignItems: 'center',
    justifyContent: 'center',
  },
  button: {
    padding:10,
    width: 100,
    backgroundColor: '#b703ff',
    borderRadius: 4,
    alignItems: 'center'
  },
  buttonText: {
    color:'white',
  }
});
