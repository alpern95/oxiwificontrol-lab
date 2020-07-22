import * as React from "react";
import { List, Edit, Create, TabbedForm,FormTab,TextInput, Datagrid, TextField, EditButton } from 'react-admin';

export const BorneList = props => (
    <List {...props}>
        <Datagrid rowClick="edit">
            <TextField source="nom" />
            <TextField source="description" /> 
            <TextField source="device" />
            <TextField source="adresse" />
            <TextField source="groupe" />
            <TextField source="modele" />
            <TextField source="username" />
            <TextField source="password" />
            <TextField source="enablepassword" />
            <TextField source="interface" />
            <TextField source="etat" />
            <TextField source="lastrefresh" />
            <EditButton />
        </Datagrid>
    </List>
);


export const BorneEdit = props => (
    <Edit {...props} >
    <TabbedForm>
      <FormTab
          label="resources.products.tabs.image"
      >
      <TextInput source="nom" />
      <TextInput source="description" />
      </FormTab>
    
        <FormTab 
        label="resources.bornes.tabs.details" 
        >
        <TextInput source="device" />
        <TextInput source="adresse" />
        <TextInput source="groupe" />
        <TextInput source="modele" />
        <TextInput source="password" />
        <TextInput source="username" />
        <TextInput source="password" />
        <TextInput source="enablepassword" />
        <TextInput source="interface" />
        <TextInput source="etat" />
        <TextInput source="lastrefresh" />
        <TextInput source="utilisateurs" />
      </FormTab>
    </TabbedForm>
  </Edit>
);

export const BorneCreate = props => (
    <Create {...props} >
    <TabbedForm>
      <FormTab
          label="resources.products.tabs.image"
      >
      <TextInput source="nom" />
      <TextInput source="description" />
      </FormTab>
        <FormTab
        label="resources.bornes.tabs.details"
        >
        <TextInput source="device" />
        <TextInput source="adresse" />
        <TextInput source="groupe" />
        <TextInput source="modele" />
        <TextInput source="password" />
        <TextInput source="modele" />
        <TextInput source="username" />
        <TextInput source="password" />
        <TextInput source="enablepassword" />
        <TextInput source="interface" />
        <TextInput source="etat" />
        <TextInput source="lastrefresh" />
        <TextInput source="utilisateurs" />
      </FormTab>
    </TabbedForm>
  </Create>
);
