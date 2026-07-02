package entities

type EvolutionAssets struct {
    Year                int     `json:"year"`
    Age                 int     `json:"age"`
    FinalAssets         float64 `json:"finalAssets"`
    InflationAdjusted   float64 `json:"finalAssetsAdjusted"`
}