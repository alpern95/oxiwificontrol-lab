Oxiwificontrol
==============
[![Build Status](https://github.com/alpern95/oxiwificontrol-lab/blob/master/build-status.png)](https://github.com/alpern95/oxiwificontrol-lab)

Un outil permettant la gestion de borne wifi dans les aglomérations pour les bibliothéque, Médiathéque

Les bornes wifi peuvent être éteintes, allumées à la demande par l'utilisateur responsable d'un groupe de bornes.

Cette outil possède:
 
* un backend écrit en go et accessible par une restfull API.
  Ce backend gére une base mongdb
  Ce backend dispose d'un outil de connexion au switch où sont connectés les bornes.

* un frontend qui est développé avec  react-admin 

### Fonctionnalités
  
- Interface d'administration
- Multi-utilsateurs
- Rôle par utilisateur
- Groupe de bornes 
- Proxy ssl (zéro conf)
- Accès personnalisables à différents modèles de Switch (EXOS)

### Architecture
![Architecture](https://github.com/alpern95/oxiwificontrol-lab/blob/master/archi.png)

 
### En France la loi indique:

Dans les établissements accueillant les enfants de moins de 3 ans, la loi interdit le WiFi dans les espaces dédiés à l'accueil, au repos et aux activités.

Dans les classes des écoles primaires où la commune a installé du WiFi, il doit être coupé lorsqu'il n'est pas utilisé pour les activités pédagogiques. Pour toute nouvelle installation, la commune doit en informer au préalable le conseil d'école.

### Ref
- le proxy ssl de [suyashkumar!](https://github.com/suyashkumar/ssl-proxy)

- le serveur http de [kevinpollet](https://github.com/kevinpollet/srv/)

- le fontend React-Admin [marmelab](https://github.com/marmelab/react-admin)

- le serveur restful api [go-restfull](github.com/emicklei/go-restful)

### SrcreenShots

![login form](https://github.com/alpern95/oxiwificontrol-lab/blob/master/image1.png)


![liste_bornes](https://github.com/alpern95/oxiwificontrol-lab/blob/master/image2.png)


![stop_bornes](https://github.com/alpern95/oxiwificontrol-lab/blob/master/image3.png)


![stop_bornes_1](https://github.com/alpern95/oxiwificontrol-lab/blob/master/image4.png)
