import {gql} from '@apollo/client';

export const ADD_PLANT = gql`
  mutation AddPlant(
    $name: String!
    $otherNames: String
    $description: String
    $plantSeason: String
    $harvestSeason: String
    $pruneSeason: String
    $tips: String
  ) {
    addPlant(
      name: $name
      otherNames: $otherNames
      description: $description
      plantSeason: $plantSeason
      harvestSeason: $harvestSeason
      pruneSeason: $pruneSeason
      tips: $tips
    ) {
      _id
      name
      otherNames
      description
      plantSeason
      harvestSeason
      pruneSeason
      tips
    }
  }
`;
