# Raspberry-Model-and-Type

# Documentation d'utilisation de l'application

## Objectif

Cette application permet de récupérer des informations spécifiques sur un Raspberry Pi, à savoir son modèle et sa quantité de mémoire. Pour cela, vous devez exécuter le programme en fournissant les paramètres appropriés pour l'utilisateur, le mot de passe, et l'hôte.

## Instructions d'utilisation

Pour exécuter l'application, vous devez fournir les trois paramètres suivants :

1. **`-u`** : Spécifie le nom d'utilisateur.
2. **`-p`** : Spécifie le mot de passe correspondant à l'utilisateur.
3. **`-h`** : Spécifie l'hôte cible, qui peut être un nom de domaine ou une adresse IP.

### Exemple de commande

```bash
application -u=test -p=pouete -h=target
application -u=test -p=pouete -h="10.10.10.10"

application --help
