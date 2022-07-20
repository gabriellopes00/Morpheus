import { MigrationInterface, QueryRunner, Table } from 'typeorm'

export class CreateLocationsMigration implements MigrationInterface {
  name = this.constructor.name.concat('1657896603908')

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.createTable(
      new Table({
        name: 'locations',
        columns: [
          { name: 'id', type: 'uuid', isPrimary: true },
          { name: 'event_id', type: 'uuid', isNullable: false },
          { name: 'street', type: 'varchar', isNullable: false },
          { name: 'district', type: 'varchar', isNullable: false },
          { name: 'state', type: 'char', length: '2', isNullable: false },
          { name: 'postal_code', type: 'char', length: '8', isNullable: false },
          { name: 'city', type: 'varchar', isNullable: false },
          { name: 'description', type: 'text', isNullable: true },
          { name: 'number', type: 'int', isNullable: false },
          { name: 'latitude', type: 'decimal', precision: 10, scale: 8, isNullable: true },
          { name: 'longitude', type: 'decimal', precision: 10, scale: 8, isNullable: true },
          { name: 'created_at', type: 'timestamp', default: 'now()' },
          { name: 'updated_at', type: 'timestamp', default: 'now()' }
        ]
      }),
      true
    )
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.dropTable('locations', true)
  }
}
