import { MigrationInterface, QueryRunner, Table, TableForeignKey } from 'typeorm'

export class CreateEventsMigration implements MigrationInterface {
  name = this.constructor.name.concat('1657896600958')

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.createTable(
      new Table({
        name: 'events',
        columns: [
          { name: 'id', type: 'uuid', isPrimary: true },
          { name: 'name', type: 'varchar', isNullable: false },
          { name: 'description', type: 'varchar', isNullable: true },
          { name: 'cover_url', type: 'varchar' },
          { name: 'organizer_account_id', type: 'uuid', isNullable: false },
          {
            name: 'age_group',
            type: 'int',
            enum: ['0', '10', '12', '14', '16', '18'],
            isNullable: false
          },
          {
            name: 'status',
            type: 'varchar',
            enum: ['available', 'finished', 'sold_out', 'canceled'],
            isNullable: false
          },
          { name: 'start_date_time', type: 'timestamp', isNullable: false },
          { name: 'end_date_time', type: 'timestamp', isNullable: false },
          { name: 'category_id', type: 'uuid', isNullable: false },
          { name: 'visibility', type: 'varchar', enum: ['private', 'public'], isNullable: false },
          { name: 'created_at', type: 'timestamp', default: 'now()' },
          { name: 'updated_at', type: 'timestamp', default: 'now()' }
        ]
      }),
      true
    )

    await queryRunner.createForeignKey(
      'events',
      new TableForeignKey({
        name: 'fk_events_organizer_account',
        referencedTableName: 'accounts',
        referencedColumnNames: ['id'],
        columnNames: ['organizer_account_id'],
        onDelete: 'cascade',
        onUpdate: 'cascade'
      })
    )

    await queryRunner.createForeignKey(
      'events',
      new TableForeignKey({
        name: 'fk_events_categories',
        referencedTableName: 'categories',
        referencedColumnNames: ['id'],
        columnNames: ['category_id'],
        onDelete: 'cascade',
        onUpdate: 'cascade'
      })
    )
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.dropForeignKey('events', 'fk_events_organizer_account')
    await queryRunner.dropForeignKey('events', 'fk_events_categories')
    await queryRunner.dropTable('events', true)
  }
}
