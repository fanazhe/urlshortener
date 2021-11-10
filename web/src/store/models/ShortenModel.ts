export default class ShortenModel {
  Id: string;
  Url: string;
  CreatedAt: number;
  Visits: number;

  constructor(id: string, url: string, createdAt: number, visits: number) {
    this.Id = id;
    this.Url = url;
    this.CreatedAt = createdAt;
    this.Visits = visits;
  }
}
